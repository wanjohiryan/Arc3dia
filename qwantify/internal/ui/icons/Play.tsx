import * as React from "react"
import { SVGProps } from "react"

const SvgComponent = (props: SVGProps<SVGSVGElement>) => (
  <svg width={24} height={24} viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" {...props}>
    <g
      stroke="#FFF"
      strokeWidth={1.5}
      fill="none"
      fillRule="evenodd"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M12 2.5c5.246 0 9.5 4.253 9.5 9.5s-4.254 9.5-9.5 9.5A9.5 9.5 0 0 1 2.5 12 9.5 9.5 0 0 1 12 2.5Z" />
      <path d="M15 11.995c0-.811-4.157-3.406-4.63-2.94-.47.467-.516 5.369 0 5.88.518.512 4.63-2.129 4.63-2.94Z" />
    </g>
  </svg>
)

export default SvgComponent
