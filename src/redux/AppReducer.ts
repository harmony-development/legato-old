import { IState, Action, Actions } from '../types/redux';
import { red, orange } from '@material-ui/core/colors';

const appState: IState = {
    theme: {
        type: 'dark',
        primary: red,
        secondary: orange
    }
};

export default function AppReducer(state = appState, action: Action): IState {
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
        default: {
            return state;
        }
    }
}
