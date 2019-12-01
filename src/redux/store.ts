import { createStore } from 'redux';
import AppReducer from './AppReducer';

export const store = createStore(AppReducer);
