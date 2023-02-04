import { random } from "../functions/utils"
import { StyledBalloon } from "./FloatingBalloon"

interface props {
    text: string
    color: string
    id: number
}

const Balloon = ({text, color, id}: props) => {
  return (
    <div>
       <StyledBalloon
        color={color}
        show={true}
        visible={true}
        animate={{
          left: random(-10, 80),
          loop: false,
          duration: random(5, 8),
          delay: 2,
          hangOnTop: false
        }}
      ><div className="string">
      </div>
      {<span className="msg ">{text}</span>}
      </StyledBalloon>
    </div>
  )
}

export default Balloon