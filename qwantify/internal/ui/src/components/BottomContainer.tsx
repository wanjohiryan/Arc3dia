import React, { RefObject, useState } from 'react';
import Info from '../../icons/Info Square'
import Settings from '../../icons/Setting'
import VolumeOn from '../../icons/Volume On'
import VolumeOff from '../../icons/Volume Off'
import FullScreen from '../../icons/Fullscreen'

interface Props {
    infoRef: RefObject<HTMLDivElement | null>;
    settingsRef: RefObject<HTMLDivElement | null>;
    playerRef: RefObject<HTMLDivElement | null>;
    vidRef: RefObject<HTMLVideoElement | null>;
}

export default function BottomContainer({ infoRef, settingsRef, playerRef, vidRef }: Props) {

    const [videoMuted, setVideoMuted] = useState(() => vidRef.current ? vidRef.current.muted : true);

    return (
        <div id="bottom-container">
            <div id="bottom-controls">
                <div
                    onClick={(e) => {
                        // close settings card and show info card
                        infoRef.current && infoRef.current.classList.remove("hide-display");
                        settingsRef.current && settingsRef.current.classList.add("hide-display");

                        e.preventDefault()
                    }}
                    id="bottom-btn">
                    <Info className='bottom-icon' height={50} width={50} />
                </div>

                <div
                    onClick={(e) => {
                        // show settings card and close info card
                        infoRef.current && infoRef.current.classList.add("hide-display");
                        settingsRef.current && settingsRef.current.classList.remove("hide-display");

                        e.preventDefault()
                    }}
                    id="bottom-btn">
                    <Settings className='bottom-icon' height={50} width={50} />
                </div>

                <div
                    onClick={(e) => {

                        setVideoMuted(e => !e);
                        if (vidRef.current)
                            vidRef.current.muted = videoMuted; //.muted = true;

                        e.preventDefault();
                    }}
                    id="bottom-btn">
                    {videoMuted ? (
                        <VolumeOn className='bottom-icon' height={50} width={50}/>
                       ) : (
                        <VolumeOff className='bottom-icon' height={50} width={50}/>
                    )}
                </div>


                <div
                    onClick={(e) => {
                        settingsRef.current && settingsRef.current.classList.add("hide-display");
                        infoRef.current && infoRef.current.classList.add("hide-display");

                        //make background fullscreen
                        playerRef.current && playerRef.current.requestFullscreen();
                    }}
                    id="bottom-btn">
                   <FullScreen className='bottom-icon' height={50} width={50}/>
                </div>
            </div>
        </div >
    )
}