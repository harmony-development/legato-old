import { configureStore, combineReducers } from '@reduxjs/toolkit';

import { AppReducer } from './AppReducer';
import { AuthReducer } from './AuthReducer';

const rootReducer = combineReducers({
	app: AppReducer,
	auth: AuthReducer,
});

export const store = configureStore({
	reducer: rootReducer,
});

export type RootState = ReturnType<typeof rootReducer>;
export type AppDispatch = typeof store.dispatch;
