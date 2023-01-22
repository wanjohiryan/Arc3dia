import React, { RefObject, useRef, useState } from 'react';
import InfoContainer from './InfoContainer';
import SettingsContainer from './SettingsContainer';
import BottomContainer from './BottomContainer';

interface Props {
    infoRef: RefObject<HTMLDivElement | null>;
    settingsRef: RefObject<HTMLDivElement | null>;
    playerRef: RefObject<HTMLDivElement | null>;
    vidRef: RefObject<HTMLVideoElement | null>;
    player:any;
}

export default function PlayScreenTwo({ infoRef, player, playerRef, vidRef}: Props) {
    const settingsRef = useRef(null);

    // on startup, check whether user prefers live or normal latency mode
    const latency:any = localStorage.getItem("live");

    const [latencyMode, setLatencyMode] = useState<string>(()=>latency ? "Low" : "Normal");

    return (
        <div id="play2">

            <InfoContainer {...{ infoRef, latencyMode }} />

            <SettingsContainer {...{ settingsRef, setLatencyMode, player }} />

            <BottomContainer {...{settingsRef, infoRef, playerRef, vidRef}}/>
        </div>
    )
}