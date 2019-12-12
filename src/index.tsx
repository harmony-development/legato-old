import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './index.css';
import './Root/ContextMenu.css';
import Root from './Root/Root';
import setupRedux from './redux/store';
import { PersistGate } from 'redux-persist/integration/react';

const { store, persistor } = setupRedux();

const ReduxRoot: React.FC = () => {
    return (
        <Provider store={store}>
            <PersistGate loading={null} persistor={persistor}>
                <Root />
            </PersistGate>
        </Provider>
    );
};

ReactDOM.render(<ReduxRoot />, document.getElementById('root'));
