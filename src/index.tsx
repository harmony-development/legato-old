import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import './index.css';
import './Root/ContextMenu.css';
import Root from './Root/Root';
import { store } from './redux/store';

const ReduxRoot: React.FC = () => {
    return (
        <Provider store={store}>
            <Root />
        </Provider>
    );
};

ReactDOM.render(<ReduxRoot />, document.getElementById('root'));
