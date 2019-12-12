import React, { useRef, useEffect } from 'react';
import { TextField } from '@material-ui/core';
import { useSelector } from 'react-redux';
import { IState } from '../../../../types/redux';
import { harmonySocket } from '../../../Root';

export const Input = () => {
    const connected = useSelector((state: IState) => state.connected);
    const inputStyle = useSelector((state: IState) => state.inputStyle);
    const guildID = useSelector((state: IState) => state.selectedGuild);
    const focus = useSelector((state: IState) => state.chatInputFocus);
    const channelID = useSelector((state: IState) => state.selectedChannel);
    const inputField = useRef<HTMLInputElement | undefined>();

    const onKeyPress = (e: React.KeyboardEvent) => {
        if (e.key === 'Enter' && !e.shiftKey) {
            e.preventDefault();
            // does the input field exist and is it not blank
            if (inputField.current && !/^\s*$/.test(inputField.current.value) && channelID) {
                harmonySocket.sendMessage(guildID, channelID, inputField.current.value);
                inputField.current.value = '';
            }
        }
    };

    useEffect(() => {
        if (inputField.current) {
            inputField.current.focus();
        }
    }, [focus]);

    return (
        <div>
            <TextField
                label={connected ? 'Message' : 'Currently Offline'}
                variant={inputStyle as any}
                fullWidth
                multiline
                rowsMax={3}
                rows={3}
                onKeyPress={onKeyPress}
                inputRef={inputField}
                color='secondary'
            />
        </div>
    );
};
