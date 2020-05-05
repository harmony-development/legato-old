import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';
import { persistStore, persistReducer, PERSIST } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import hardSet from 'redux-persist/es/stateReconciler/hardSet';

import { IState } from '../types/redux';

import { AppReducer } from './AppReducer';
const persisted = persistReducer<IState>(
	{
		key: 'root',
		storage,
		stateReconciler: hardSet,
	},
	AppReducer
);

export const store = configureStore({
	reducer: persisted,
	middleware: getDefaultMiddleware({
		serializableCheck: {
			ignoredActions: [PERSIST],
		},
	}),
});

export const persistor = persistStore(store);

export type AppDispatch = typeof store.dispatch;
