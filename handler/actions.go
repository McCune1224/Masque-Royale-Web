package handler

import (
	"log"
	"math"
	"math/rand"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/views"
)

// Key value of queue order of action types
var (
	ActionPriority = map[string]int{
		"Alteration":                           1,
		"Reactive":                             2,
		"Redirection":                          3,
		"Visit Redirection":                    3,
		"Investigation":                        4,
		"Protection":                           5,
		"Visit Blocking":                       6,
		"Vote Immunity":                        7,
		"Vote Manipulation":                    8,
		"Vote Blocking":                        8,
		"Support":                              9,
		"Debuff":                               10,
		"Theft":                                11,
		"Healing":                              12,
		"Destruction":                          13,
		"Killing":                              14,
		"Last: (exception): Detective's Solve": 15,
	}
)

func insert[T any](items []T, item T, i int) {
	items = append(items[:i+1], items[i:]...)
	items[i] = item
}

func scramble[T any](items []T) []T {
	perms := rand.Perm(len(items))
	newItems := make([]T, len(items))
	for i := 0; i < len(items); i++ {
		p := perms[i]
		newItems[i] = items[p]
	}

	return newItems
}

func highestPriority(action data.Action) int {
	lowest := math.MaxInt
	for _, cat := range action.Categories {
		rank := ActionPriority[cat]
		if lowest > rank && rank != 0 {
			// have to check for 0 because of default value in maps
			lowest = rank
		}
	}
	return lowest
}

func flatten[T any](lists [][]T) []T {
	var res []T
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func shuffle[T any](items []T) []T {
	perms := rand.Perm(len(items))
	newItems := make([]T, len(items))
	for i := 0; i < len(items); i++ {
		p := perms[i]
		newItems[i] = items[p]
	}
	return newItems
}

func sortActionByPriority(actions []data.Action) []data.Action {
	buckets := make([][]data.Action, len(ActionPriority))

	for _, v := range actions {
		prio := highestPriority(v)
		buckets[prio] = append(buckets[prio], v)
	}

	for i := range buckets {
		buckets[i] = shuffle(buckets[i])
	}

	return flatten(buckets)
}

func (h *Handler) ActionDashboard(c echo.Context) error {
	actions, err := h.models.Actions.GetAll()
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	gameActionList, err := h.models.Actions.GetActionList(c.Param("game_id"))
	if err != nil {
		log.Print(err)
		c.Redirect(302, "/")
	}
	var queue []data.Action

	for _, id := range gameActionList.ActionIds {
		q, err := h.models.Actions.GetByID(int(id))
		if err != nil {
			log.Println(err)
			return c.Redirect(302, "/")
		}
		queue = append(queue, *q)
	}

	foo := sortActionByPriority(queue)

	return TemplRender(c, views.ActionDashboard(c, actions, foo))
}

func (h *Handler) UpdateActionsListItem(c echo.Context) error {
	gameActionList, err := h.models.Actions.GetActionList(c.Param("game_id"))
	if err != nil {
		log.Print(err)
		return c.Redirect(302, "/")
	}
	actionForm := c.FormValue("action")
	newAction, err := h.models.Actions.GetByName(actionForm)
	if err != nil {
		log.Print(err)
		return c.Redirect(302, "/")
	}

	gameActionList.ActionIds = append(gameActionList.ActionIds, int64(newAction.ID))
	err = h.models.Actions.UpdateActionList(gameActionList)
	if err != nil {
		log.Print(err)
		return c.Redirect(302, "/")
	}
	log.Println("-------", gameActionList.ActionIds)

	var actions []data.Action
	for _, id := range gameActionList.ActionIds {
		action, err := h.models.Actions.GetByID(int(id))
		if err != nil {
			log.Print(err)
			return c.Redirect(302, "/")
		}
		actions = append(actions, *action)
	}

	actions = sortActionByPriority(actions)
	return TemplRender(c, views.ActionQueue(c, actions))
}

func (h *Handler) RemoveActionListItem(c echo.Context) error {
	gameActionList, err := h.models.Actions.GetActionList(c.Param("game_id"))
	if err != nil {
		log.Print(err)
		return c.Redirect(302, "/")
	}

	actionForm := c.QueryParam("name")
	actionID, _ := strconv.Atoi(actionForm)

	for i, id := range gameActionList.ActionIds {
		if int(id) == actionID {
			// getting panic when there are 2 Ids with same value in the list to fix this:
			// we can use a break statement to exit the loop after deleting the first instance of the id
			if len(gameActionList.ActionIds) == 2 && gameActionList.ActionIds[0] == gameActionList.ActionIds[1] {
				gameActionList.ActionIds = pq.Int64Array{gameActionList.ActionIds[0]}
				break
			}
			// safe delete from slice (in case of duplicate names, we only want to delete the first instance)
			gameActionList.ActionIds = append(gameActionList.ActionIds[:i], gameActionList.ActionIds[i+1:]...)
		}
	}

	err = h.models.Actions.UpdateActionList(gameActionList)
	if err != nil {
		log.Print(err)
		return c.Redirect(302, "/")
	}
	log.Println("-------", gameActionList.ActionIds)

	var actions []data.Action
	for _, id := range gameActionList.ActionIds {
		action, err := h.models.Actions.GetByID(int(id))
		if err != nil {
			log.Print(err)
			return c.Redirect(302, "/")
		}
		actions = append(actions, *action)
	}

	actions = sortActionByPriority(actions)
	return TemplRender(c, views.ActionQueue(c, actions))
}
