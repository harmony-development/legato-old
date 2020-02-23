import { configureStore } from '@reduxjs/toolkit';

import { AppReducer } from './AppReducer';

export const store = configureStore({
	reducer: AppReducer,
});

export type AppDispatch = typeof store.dispatch;
