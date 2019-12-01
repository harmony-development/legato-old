import React from 'react';
import { Dialog, DialogTitle, DialogContent, CardActionArea } from '@material-ui/core';

const ImageDialog = (props) => {
    return (
        <Dialog open={props.open} onClose={() => props.setOpen(false)}>
            <DialogTitle>Image Preview</DialogTitle>
            <DialogContent>
                <CardActionArea>
                    <img src={props.image} style={{ width: '100%' }} alt='' />
                </CardActionArea>
            </DialogContent>
        </Dialog>
    );
};

export default ImageDialog;
