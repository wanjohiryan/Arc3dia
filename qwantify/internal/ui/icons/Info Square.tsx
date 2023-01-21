import * as React from "react"
import { SVGProps } from "react"

const SvgComponent = (props: SVGProps<SVGSVGElement>) => (
  <svg width={24} height={24} viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" {...props}>
    <g
      stroke="#FFF"
      fill="none"
      fillRule="evenodd"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path
        d="M16.334 2.75H7.665c-3.02 0-4.915 2.14-4.915 5.166v8.168c0 3.027 1.885 5.166 4.915 5.166h8.668c3.031 0 4.917-2.139 4.917-5.166V7.916c0-3.027-1.886-5.166-4.916-5.166ZM11.995 16v-4"
        strokeWidth={1.5}
      />
      <path strokeWidth={2} d="M11.99 8.204H12" />
    </g>
  </svg>
)

export default SvgComponent
