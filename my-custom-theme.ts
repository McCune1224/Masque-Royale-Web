import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const myCustomTheme: CustomThemeConfig = {
	name: 'my-custom-theme',
	properties: {
		// =~= Theme Properties =~=
		'--theme-font-family-base': `system-ui`,
		'--theme-font-family-heading': `system-ui`,
		'--theme-font-color-base': '0 0 0',
		'--theme-font-color-dark': '255 255 255',
		'--theme-rounded-base': '9999px',
		'--theme-rounded-container': '8px',
		'--theme-border-base': '1px',
		// =~= Theme On-X Colors =~=
		'--on-primary': '0 0 0',
		'--on-secondary': '0 0 0',
		'--on-tertiary': '0 0 0',
		'--on-success': '255 255 255',
		'--on-warning': '255 255 255',
		'--on-error': '0 0 0',
		'--on-surface': '255 255 255',
		// =~= Theme Colors  =~=
		// primary | #c0639d
		'--color-primary-50': '246 232 240', // #f6e8f0
		'--color-primary-100': '242 224 235', // #f2e0eb
		'--color-primary-200': '239 216 231', // #efd8e7
		'--color-primary-300': '230 193 216', // #e6c1d8
		'--color-primary-400': '211 146 186', // #d392ba
		'--color-primary-500': '192 99 157', // #c0639d
		'--color-primary-600': '173 89 141', // #ad598d
		'--color-primary-700': '144 74 118', // #904a76
		'--color-primary-800': '115 59 94', // #733b5e
		'--color-primary-900': '94 49 77', // #5e314d
		// secondary | #ec44ba
		'--color-secondary-50': '252 227 245', // #fce3f5
		'--color-secondary-100': '251 218 241', // #fbdaf1
		'--color-secondary-200': '250 208 238', // #fad0ee
		'--color-secondary-300': '247 180 227', // #f7b4e3
		'--color-secondary-400': '242 124 207', // #f27ccf
		'--color-secondary-500': '236 68 186', // #ec44ba
		'--color-secondary-600': '212 61 167', // #d43da7
		'--color-secondary-700': '177 51 140', // #b1338c
		'--color-secondary-800': '142 41 112', // #8e2970
		'--color-secondary-900': '116 33 91', // #74215b
		// tertiary | #78e2e3
		'--color-tertiary-50': '235 251 251', // #ebfbfb
		'--color-tertiary-100': '228 249 249', // #e4f9f9
		'--color-tertiary-200': '221 248 248', // #ddf8f8
		'--color-tertiary-300': '201 243 244', // #c9f3f4
		'--color-tertiary-400': '161 235 235', // #a1ebeb
		'--color-tertiary-500': '120 226 227', // #78e2e3
		'--color-tertiary-600': '108 203 204', // #6ccbcc
		'--color-tertiary-700': '90 170 170', // #5aaaaa
		'--color-tertiary-800': '72 136 136', // #488888
		'--color-tertiary-900': '59 111 111', // #3b6f6f
		// success | #b9322b
		'--color-success-50': '245 224 223', // #f5e0df
		'--color-success-100': '241 214 213', // #f1d6d5
		'--color-success-200': '238 204 202', // #eeccca
		'--color-success-300': '227 173 170', // #e3adaa
		'--color-success-400': '206 112 107', // #ce706b
		'--color-success-500': '185 50 43', // #b9322b
		'--color-success-600': '167 45 39', // #a72d27
		'--color-success-700': '139 38 32', // #8b2620
		'--color-success-800': '111 30 26', // #6f1e1a
		'--color-success-900': '91 25 21', // #5b1915
		// warning | #0057e4
		'--color-warning-50': '217 230 251', // #d9e6fb
		'--color-warning-100': '204 221 250', // #ccddfa
		'--color-warning-200': '191 213 248', // #bfd5f8
		'--color-warning-300': '153 188 244', // #99bcf4
		'--color-warning-400': '77 137 236', // #4d89ec
		'--color-warning-500': '0 87 228', // #0057e4
		'--color-warning-600': '0 78 205', // #004ecd
		'--color-warning-700': '0 65 171', // #0041ab
		'--color-warning-800': '0 52 137', // #003489
		'--color-warning-900': '0 43 112', // #002b70
		// error | #258789
		'--color-error-50': '222 237 237', // #deeded
		'--color-error-100': '211 231 231', // #d3e7e7
		'--color-error-200': '201 225 226', // #c9e1e2
		'--color-error-300': '168 207 208', // #a8cfd0
		'--color-error-400': '102 171 172', // #66abac
		'--color-error-500': '37 135 137', // #258789
		'--color-error-600': '33 122 123', // #217a7b
		'--color-error-700': '28 101 103', // #1c6567
		'--color-error-800': '22 81 82', // #165152
		'--color-error-900': '18 66 67', // #124243
		// surface | #45749d
		'--color-surface-50': '227 234 240', // #e3eaf0
		'--color-surface-100': '218 227 235', // #dae3eb
		'--color-surface-200': '209 220 231', // #d1dce7
		'--color-surface-300': '181 199 216', // #b5c7d8
		'--color-surface-400': '125 158 186', // #7d9eba
		'--color-surface-500': '69 116 157', // #45749d
		'--color-surface-600': '62 104 141', // #3e688d
		'--color-surface-700': '52 87 118', // #345776
		'--color-surface-800': '41 70 94', // #29465e
		'--color-surface-900': '34 57 77' // #22394d
	}
};
