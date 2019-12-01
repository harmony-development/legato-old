import React from 'react';
import { HarmonyBar } from './HarmonyBar/HarmonyBar';
import { ThemeDialog } from './Dialog/ThemeDialog';

export const App = () => {
    return (
        <div>
            <ThemeDialog />
            <HarmonyBar />
        </div>
    );
};
