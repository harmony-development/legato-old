import React from 'react';
import { ButtonBase, Color, Typography } from '@material-ui/core';
import { red, yellow, orange, blue, green, indigo } from '@material-ui/core/colors';
import CheckCircle from '@material-ui/icons/CheckCircle';

interface IPickerState {
    color: Color;
    setColor: React.Dispatch<React.SetStateAction<Color>>;
    label: string;
}

export const ColorPicker = (state: IPickerState) => {
    return (
        <div>
            <Typography>{state.label}</Typography>
            <div style={{ display: 'flex' }}>
                <ButtonBase style={{ backgroundColor: red[500], width: '40px', height: '40px' }} onClick={() => state.setColor(red)}>
                    {state.color === red ? <CheckCircle /> : undefined}
                </ButtonBase>
                <ButtonBase style={{ backgroundColor: orange[500], width: '40px', height: '40px' }} onClick={() => state.setColor(orange)}>
                    {state.color === orange ? <CheckCircle /> : undefined}
                </ButtonBase>
                <ButtonBase style={{ backgroundColor: yellow[500], width: '40px', height: '40px' }} onClick={() => state.setColor(yellow)}>
                    {state.color === yellow ? <CheckCircle /> : undefined}
                </ButtonBase>
                <ButtonBase style={{ backgroundColor: green[500], width: '40px', height: '40px' }} onClick={() => state.setColor(green)}>
                    {state.color === green ? <CheckCircle /> : undefined}
                </ButtonBase>
                <ButtonBase style={{ backgroundColor: blue[500], width: '40px', height: '40px' }} onClick={() => state.setColor(blue)}>
                    {state.color === blue ? <CheckCircle /> : undefined}
                </ButtonBase>
                <ButtonBase style={{ backgroundColor: indigo[500], width: '40px', height: '40px' }} onClick={() => state.setColor(indigo)}>
                    {state.color === indigo ? <CheckCircle /> : undefined}
                </ButtonBase>
            </div>
        </div>
    );
};
