import React, { RefObject, useState } from 'react';
import Close from "../../icons/Close"

interface Props {
    infoRef: RefObject<HTMLDivElement | null>;
    latencyMode: string;
}

export default function InfoContainer({ infoRef, latencyMode }: Props) {

    return (
        <>

            <div
                //@ts-expect-error
                ref={infoRef} id="info-container" className="info-container hide-display">
                <div id="close-info-tab">
                    <div
                        id="close-btn"
                        onClick={(e) => {
                            // close info card
                            infoRef.current && infoRef.current.classList.add("hide-display");

                            e.preventDefault()
                        }}>
                            <Close className="close-icon" height={15} width={15} />
                    </div>
                </div>

                <div id="name-value-container">
                    <div id="duo-containers">Name</div>
                    <div id="duo-containers">Value</div>
                </div>

                <div id="name-value-container" className="video-codec">
                    <div className="text-style" >Video Codec</div>
                    <div id="video-codec" className="text-style"></div>
                </div>

                <div id="name-value-container" className="audio-codec">
                    <div className="text-style">Audio Codec</div>
                    <div id="audio-codec" className="text-style"></div>
                </div>

                <div id="name-value-container" className="video-buffer">
                    <div className="text-style">Video Buffer</div>
                    <div id="video-buffer" className="buffer-style"></div>
                </div>

                <div id="name-value-container" className="audio-buffer">
                    <div className="text-style" >Audio Buffer</div>
                    <div id="audio-buffer" className="buffer-style"></div>
                </div>

                <div id="name-value-container" className="latency-source">
                    <div className="text-style">Latency from source</div>
                    <div id="latency-source" className="text-style"></div>
                </div>

                <div id="name-value-container" className="vid-res" >
                    <div className="text-style">Video resolution</div>
                    <div id="vid-res" className="text-style"></div>
                </div>

                <div id="name-value-container" className="latency-mode">
                    <div className="text-style">Latency mode</div>
                    <div id="latency-mode" className="text-style">{latencyMode}</div>
                </div>

                {/* <div id="name-value-container">
                    <div className="text-style">Protocol</div>
                    <div className="text-style">HLS</div>
            </div> */}


            </div>
        </>
    )
}