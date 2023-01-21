import * as React from "react"
import { SVGProps } from "react"

const SvgComponent = (props: SVGProps<SVGSVGElement>) => (
  <svg width={24} height={24} viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" {...props}>
    <path
      style={{
        stroke: "none",
        strokeWidth: 1,
        strokeDasharray: "none",
        strokeLinecap: "butt",
        strokeDashoffset: 0,
        strokeLinejoin: "miter",
        strokeMiterlimit: 4,
        fill: "#fff",
        fillRule: "nonzero",
        opacity: 1,
      }}
      transform="matrix(.48 0 0 .48 .48 .48)"
      d="M8.5 7C5.48 7 3 9.48 3 12.5v23C3 38.52 5.48 41 8.5 41h31c3.02 0 5.5-2.48 5.5-5.5v-23C45 9.48 42.52 7 39.5 7h-31zm0 3h31c1.398 0 2.5 1.102 2.5 2.5v23c0 1.398-1.102 2.5-2.5 2.5H28v-8.5a5.5 5.5 0 0 0-5.5-5.5H6V12.5C6 11.102 7.102 10 8.5 10zm28.97 2.986a1.5 1.5 0 0 0-.16.014H32.5a1.5 1.5 0 1 0 0 3h1.379l-5.44 5.44a1.5 1.5 0 1 0 2.122 2.12L36 18.122V19.5a1.5 1.5 0 1 0 3 0v-4.826a1.5 1.5 0 0 0-1.53-1.688z"
    />
  </svg>
)

export default SvgComponent
