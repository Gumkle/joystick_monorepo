// export const ARROW_UP = 'ARROW_UP'
// export const ARROW_DOWN = 'ARROW_DOWN'
// export const ARROW_RIGHT = 'ARROW_RIGHT'
// export const ARROW_LEFT = 'ARROW_LEFT'
// export const ACTION_BUTTON_1 = 'ACTION_BUTTON_1'
// export const ACTION_BUTTON_2 = 'ACTION_BUTTON_2'
// export const ACTION_BUTTON_3 = 'ACTION_BUTTON_3'
// export const ACTION_BUTTON_4 = 'ACTION_BUTTON_4'

// export const KEY_UP = 'KEY_UP'
// export const KEY_DOWN = 'KEY_DOWN'

import { keyCodes, keyStates as kStates } from '../utils/keyCodes.js'

export const keyActions = {}

for (const [key] of keyCodes) {
  keyActions[key] = key
}

for (const [key] of kStates) {
  keyActions[key] = key
}
