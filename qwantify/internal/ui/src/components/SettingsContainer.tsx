import React, { RefObject, SetStateAction, useEffect, useRef } from "react";
import Close from "../../icons/Close"

interface Props {
    settingsRef: RefObject<HTMLDivElement | null>;
    setLatencyMode: React.Dispatch<React.SetStateAction<string>>;
    player: any;
}

export default function SettingsContainer({ settingsRef, setLatencyMode, player }: Props) {

    const switchRef = useRef<HTMLInputElement>(null);

    useEffect(() => {
        if (!switchRef.current) {
            return;
        }

        switchRef.current.addEventListener("change", function (event: any) {
            let runForever;

            //TODO: tighten this check? Too loose
            if (event.target.checked) {
                //@ts-expect-error
                localStorage.setItem("live", true);
                setLatencyMode("Low")
                //TODO: call go live if buffer goes beyond a certain threshold
                // player.goLive();
            } else {
                //@ts-expect-error
                localStorage.setItem("live", false);
                setLatencyMode("Normal")
            }
        });

        return () => {
            if (switchRef.current)
                switchRef.current.removeEventListener("change", function (event: any) {
                    let runForever;

                    //TODO: tighten this check? Too loose
                    if (event.target.checked) {
                        //@ts-expect-error
                        localStorage.setItem("live", true);
                        setLatencyMode("Low")
                        //TODO: call go live if buffer goes beyond a certain threshold
                        player.goLive();
                    } else {
                        //@ts-expect-error
                        localStorage.setItem("live", false);
                        setLatencyMode("Normal")
                    }
                });
        }
    }, [])

    return (
        <div
            //@ts-expect-error
            ref={settingsRef} id="settings-container" className="settings-container hide-display">

            <div id="close-settings-tab">
                <div
                    onClick={(e) => {
                        // close settings card
                        settingsRef.current && settingsRef.current.classList.add("hide-display");

                        e.preventDefault()
                    }}
                    id="close-btn">
                    <Close className='close-icon' height={15} width={15} />
                </div>
            </div>

            <div id="latency-setting">
                <div className="settings-text-style">Lowest latency</div>
                <div className="radio-latency-style">
                    <label className="switch">
                        <input ref={switchRef} id="switch" type="checkbox" />
                        <span className="slider"></span>
                    </label>
                </div>
            </div>

        </div>
    )
}