import { Link } from "react-router-dom";
import LogoImage from "../assets/Logo.png";

function Logo() {
  return (
    <>
      <div className="h-1/5 ml-80 pt-10">
        <div className="ml-9">
          <Link to="/">
            <img src={LogoImage} alt="" />
          </Link>
        </div>
      </div>
    </>
  );
}

export default Logo;
