import React, { useRef } from 'react';
import { Avatar, Button } from '@material-ui/core';
import { useStyles } from './styles';

const UserIcon: = (props) => {
    const inputFile = useRef(null);
    const classes = useStyles();

    const onFileSelected = (event) => {
        if (event.currentTarget.files && event.currentTarget.files.length > 0) {
            const newIcon = event.currentTarget.files[0];
            if (newIcon.type.startsWith('image/') && newIcon.size < 33554432) {
                const imageReader = new FileReader();
                imageReader.readAsDataURL(newIcon);
                imageReader.addEventListener('load', () => {
                    if (typeof imageReader.result === 'string') {
                        props.setIcon(imageReader.result);
                    }
                });
            }
        }
    };

    const onIconButtonPressed = () => {
        if (inputFile.current) {
            inputFile.current.click();
        }
    };

    return (
        <div className={classes.iconRoot}>
            <Avatar src={props.icon} />
            <input type='file' ref={inputFile} style={{ display: 'none' }} onChange={onFileSelected} />
            <Button variant='contained' color='primary' className={classes.changeIconButton} onClick={onIconButtonPressed}>
                Change Icon
            </Button>
        </div>
    );
};

export default UserIcon;
