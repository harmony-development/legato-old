/**
 * This file is the controller for Redux.
 *
 * It handles dispatches from ../actions/AppActions.ts and returns a new state.
 *
 * The initial state is never mutated. It always returns a copy of the original state.
 */

import { Actions, ActionTypes, IAppState } from '../types';
import { red, purple } from '@material-ui/core/colors';

const initialState: IAppState = {
  theme: {
    type: 'dark',
    primary: red,
    secondary: purple
  },
  nameDialog: false,
  user: {
    name: `Anonymous User #${Math.floor(Math.random() * 1000)}`,
    icon: ''
  }
};

export default function AppReducer(state = initialState, action: ActionTypes): IAppState {
  switch (action.type) {
    case Actions.INVERT_THEME: {
      return {
        ...state,
        theme: {
          ...state.theme,
          type: state.theme.type === 'dark' ? 'light' : 'dark'
        }
      };
    }
    case Actions.TOGGLE_NAME_DIALOG: {
      return {
        ...state,
        nameDialog: !state.nameDialog
      };
    }
    case Actions.UPDATE_USER: {
      return {
        ...state,
        user: action.payload
      };
    }
    default:
      return state;
  }
}
