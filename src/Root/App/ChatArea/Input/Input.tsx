import React, { useRef } from 'react';
import { TextField } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';

export const Input = () => {
    const inputField = useRef<HTMLInputElement | undefined>(undefined);
    const connected = useSelector((state: IState) => state.connected);

    const onKeyPress = (e: React.KeyboardEvent) => {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            if (inputField.current && !/^\s*$/.test(inputField.current.value)) {
                // does the input field exist and is it not blank
                inputField.current.value = '';
                console.log('Send Message Here');
            }
        }
    };

    return (
        <div>
            <TextField label={connected ? 'Message' : 'Currently Offline'} variant='filled' fullWidth multiline rowsMax={3} rows={3} onKeyPress={onKeyPress} inputRef={inputField} />
        </div>
    );
};
