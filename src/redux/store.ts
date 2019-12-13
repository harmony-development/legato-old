import { createStore } from 'redux';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import AppReducer from './AppReducer';

const persistConfig = {
    key: 'root',
    storage,
    whitelist: ['theme']
};

const persistedReducer = persistReducer(persistConfig, AppReducer);

export default () => {
    const store = createStore(persistedReducer);
    const persistor = persistStore(store);
    return { store, persistor };
};
