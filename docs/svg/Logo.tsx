import * as React from "react"
import { SVGProps } from "react"
const SvgComponent = (props: SVGProps<SVGSVGElement>) => (
  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 320" {...props}>
    <g
      style={{
        fill: "currentColor",
      }}
    >
      <path
        d="M147.224 263.615q11.9 1.4 18-3l71-41 6.5-6.5 4-8v-14q-3.2-9.3-10.5-14.5-2.2-3.8-9-3l-22.5 13.5 16 10-1.5 3.5-59 34-8.5 3v-63.5l-4.5-6.5-20.5-11v86.5q2.1 9.4 8.5 14.5z"
        opacity={0.98}
        style={{
          fill: "currentColor",
          fillOpacity: 1
        }}
      />
      <path
        d="M91.224 231.615q16.7 1.7 23.5-6.5l-.5-28.5q-7.3 6.5-17.5 9v-78.5l2.5-.5 46 27 8 4h6l21.5-12.5-74.5-43.5q-5.5-2.5-15-1-9.9 3.1-15.5 10.5l-4 9v92q2.4 9.6 9.5 14.5z"
        opacity={0.98}
        style={{
          fill: "currentColor",
          fillOpacity:1
        }}
      />
      <path
        d="m162.724 196.615 6.5-2 64-37 9.5-8.5q6.2-6.3 4-21-2.1-9.9-9.5-14.5l-72-42q-6.1-4.4-18-3-10.1 2.4-15.5 9.5l-5 11v10l3.5 3.5 19 11h2.5v-18.5l3.5-.5 65.5 38.5-1.5 3.5-48 27-6.5 5.5-1 2v22z"
        opacity={0.98}
        style={{
          fill: "currentColor",
          fillOpacity: 1,
        }}
      />
    </g>
  </svg>
)
export default SvgComponent
