import { Link } from "react-router-dom";
import LogoImage from "../assets/Logo.png";
import { userInfo } from "../stores/useStore";

function Logo() {
  const { username } = userInfo();
  return (
    <>
      <div className="h-1/5 ml-[410px] pt-10">
        <div className="flex flex-row ml-9 place-content-between mr-10">
          <Link to="/">
            <img src={LogoImage} alt="Logo" />
          </Link>
          {username ? (
            <div className="place-self-center justify-self-end">
              <p className="text-xl font-semibold">Hi, {username}</p>
            </div>
          ) : (
            <></>
          )}
        </div>
      </div>
    </>
  );
}

export default Logo;
