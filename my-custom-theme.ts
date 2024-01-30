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
		'--on-primary': '255 255 255',
		'--on-secondary': '255 255 255',
		'--on-tertiary': '0 0 0',
		'--on-success': '0 0 0',
		'--on-warning': '255 255 255',
		'--on-error': '0 0 0',
		'--on-surface': '255 255 255',
		// =~= Theme Colors  =~=
		// primary | #882039
		'--color-primary-50': '237 222 225', // #eddee1
		'--color-primary-100': '231 210 215', // #e7d2d7
		'--color-primary-200': '225 199 206', // #e1c7ce
		'--color-primary-300': '207 166 176', // #cfa6b0
		'--color-primary-400': '172 99 116', // #ac6374
		'--color-primary-500': '136 32 57', // #882039
		'--color-primary-600': '122 29 51', // #7a1d33
		'--color-primary-700': '102 24 43', // #66182b
		'--color-primary-800': '82 19 34', // #521322
		'--color-primary-900': '67 16 28', // #43101c
		// secondary | #9a5177
		'--color-secondary-50': '240 229 235', // #f0e5eb
		'--color-secondary-100': '235 220 228', // #ebdce4
		'--color-secondary-200': '230 212 221', // #e6d4dd
		'--color-secondary-300': '215 185 201', // #d7b9c9
		'--color-secondary-400': '184 133 160', // #b885a0
		'--color-secondary-500': '154 81 119', // #9a5177
		'--color-secondary-600': '139 73 107', // #8b496b
		'--color-secondary-700': '116 61 89', // #743d59
		'--color-secondary-800': '92 49 71', // #5c3147
		'--color-secondary-900': '75 40 58', // #4b283a
		// tertiary | #c9a50c
		'--color-tertiary-50': '247 242 219', // #f7f2db
		'--color-tertiary-100': '244 237 206', // #f4edce
		'--color-tertiary-200': '242 233 194', // #f2e9c2
		'--color-tertiary-300': '233 219 158', // #e9db9e
		'--color-tertiary-400': '217 192 85', // #d9c055
		'--color-tertiary-500': '201 165 12', // #c9a50c
		'--color-tertiary-600': '181 149 11', // #b5950b
		'--color-tertiary-700': '151 124 9', // #977c09
		'--color-tertiary-800': '121 99 7', // #796307
		'--color-tertiary-900': '98 81 6', // #625106
		// success | #e80cb0
		'--color-success-50': '252 219 243', // #fcdbf3
		'--color-success-100': '250 206 239', // #faceef
		'--color-success-200': '249 194 235', // #f9c2eb
		'--color-success-300': '246 158 223', // #f69edf
		'--color-success-400': '239 85 200', // #ef55c8
		'--color-success-500': '232 12 176', // #e80cb0
		'--color-success-600': '209 11 158', // #d10b9e
		'--color-success-700': '174 9 132', // #ae0984
		'--color-success-800': '139 7 106', // #8b076a
		'--color-success-900': '114 6 86', // #720656
		// warning | #a0044d
		'--color-warning-50': '241 217 228', // #f1d9e4
		'--color-warning-100': '236 205 219', // #eccddb
		'--color-warning-200': '231 192 211', // #e7c0d3
		'--color-warning-300': '217 155 184', // #d99bb8
		'--color-warning-400': '189 79 130', // #bd4f82
		'--color-warning-500': '160 4 77', // #a0044d
		'--color-warning-600': '144 4 69', // #900445
		'--color-warning-700': '120 3 58', // #78033a
		'--color-warning-800': '96 2 46', // #60022e
		'--color-warning-900': '78 2 38', // #4e0226
		// error | #a08828
		'--color-error-50': '241 237 223', // #f1eddf
		'--color-error-100': '236 231 212', // #ece7d4
		'--color-error-200': '231 225 201', // #e7e1c9
		'--color-error-300': '217 207 169', // #d9cfa9
		'--color-error-400': '189 172 105', // #bdac69
		'--color-error-500': '160 136 40', // #a08828
		'--color-error-600': '144 122 36', // #907a24
		'--color-error-700': '120 102 30', // #78661e
		'--color-error-800': '96 82 24', // #605218
		'--color-error-900': '78 67 20', // #4e4314
		// surface | #383c5c
		'--color-surface-50': '225 226 231', // #e1e2e7
		'--color-surface-100': '215 216 222', // #d7d8de
		'--color-surface-200': '205 206 214', // #cdced6
		'--color-surface-300': '175 177 190', // #afb1be
		'--color-surface-400': '116 119 141', // #74778d
		'--color-surface-500': '56 60 92', // #383c5c
		'--color-surface-600': '50 54 83', // #323653
		'--color-surface-700': '42 45 69', // #2a2d45
		'--color-surface-800': '34 36 55', // #222437
		'--color-surface-900': '27 29 45' // #1b1d2d
	}
};
