import { createAction, createReducer } from '@reduxjs/toolkit';

import { IAuthState, IInstanceList } from '../types/redux';

const authState: IAuthState = {
	instanceList: {},
};

function WithPayload<T>() {
	return (t: T) => ({ payload: t });
}

export const SetInstanceList = createAction('SET_INSTANCE_LIST', WithPayload<IInstanceList>());

export const AuthReducer = createReducer(authState, builder =>
	builder.addCase(SetInstanceList, (state, action) => ({
		...state,
		instanceList: action.payload,
	}))
);
