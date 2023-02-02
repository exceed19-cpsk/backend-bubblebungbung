import { random } from "../functions/utils"
import { StyledBalloon } from "./FloatingBalloon"

interface props {
    text: string
    color: string
}

const Balloon = ({text, color}: props) => {
  return (
    <div>
       <StyledBalloon
        color={color}
        show={true}
        visible={true}
        animate={{
          left: random(-10, 90),
          loop: true,
          duration: 5,
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