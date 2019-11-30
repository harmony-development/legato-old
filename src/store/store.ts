import { createStore } from 'redux';
import AppReducer from './reducers/AppReducer';

export const store = createStore(AppReducer);
