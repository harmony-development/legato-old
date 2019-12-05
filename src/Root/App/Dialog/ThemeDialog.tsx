import React, { useState, useEffect } from 'react';
import { Dialog, DialogTitle, DialogContent, DialogActions, Button, Color, FormControlLabel, FormControl, FormLabel, RadioGroup, Radio, Typography, Select, MenuItem } from '@material-ui/core';
import { useSelector, useDispatch } from 'react-redux';
import { IState, Actions } from '../../../types/redux';
import { ColorPicker } from './ColorPicker';
import { orange, red } from '@material-ui/core/colors';

export const ThemeDialog = () => {
    const open = useSelector((state: IState) => state.themeDialog);
    const themeType = useSelector((state: IState) => state.theme.type);
    const [primary, setPrimary] = useState<Color>(red);
    const [secondary, setSecondary] = useState<Color>(orange);
    const inputStyle = useSelector((state: IState) => state.inputStyle);
    const dispatch = useDispatch();

    const changeInputStyle = (ev: React.ChangeEvent<{ value: unknown }>) => {
        if (typeof ev.target.value === 'string') {
            dispatch({ type: Actions.SET_INPUT_STYLE, payload: ev.target.value });
        }
    };

    useEffect(() => {
        dispatch({ type: Actions.CHANGE_PRIMARY, payload: primary });
    }, [primary, dispatch]);
    useEffect(() => {
        dispatch({ type: Actions.CHANGE_SECONDARY, payload: secondary });
    }, [secondary, dispatch]);

    return (
        <Dialog open={open} onClose={() => dispatch({ type: Actions.TOGGLE_THEME_DIALOG })}>
            <DialogTitle>Customize Theme</DialogTitle>
            <DialogContent>
                <FormControl component='fieldset'>
                    <FormLabel component='legend'>Theme Type</FormLabel>
                    <RadioGroup value={themeType} row onChange={(e: React.ChangeEvent<HTMLInputElement>) => dispatch({ type: Actions.INVERT_THEME })}>
                        <FormControlLabel value='light' control={<Radio color='secondary' />} label='Light' labelPlacement='end' />
                        <FormControlLabel value='dark' control={<Radio color='secondary' />} label='Dark' labelPlacement='end' />
                    </RadioGroup>
                </FormControl>
                <ColorPicker color={primary} setColor={setPrimary} label={'Primary Color'} />
                <ColorPicker color={secondary} setColor={setSecondary} label={'Secondary Color'} />
                <Typography>Text Input Style</Typography>
                <Select value={inputStyle || 'standard'} onChange={changeInputStyle} variant={inputStyle as any} fullWidth>
                    <MenuItem value={'standard'}>Standard</MenuItem>
                    <MenuItem value={'filled'}>Filled</MenuItem>
                    <MenuItem value={'outlined'}>Outlined</MenuItem>
                </Select>
            </DialogContent>
            <DialogActions>
                <Button color='primary' onClick={() => dispatch({ type: Actions.TOGGLE_THEME_DIALOG })}>
                    Close
                </Button>
            </DialogActions>
        </Dialog>
    );
};
