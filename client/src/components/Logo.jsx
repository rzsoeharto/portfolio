import { Link } from "react-router-dom";
import LogoImage from "../assets/Logo.png";
import { userInfo } from "../stores/useStore";

function Logo() {
  const { username } = userInfo();
  return (
    <>
      <div className="flex h-40 w-full self-end">
        <div className="flex flex-row place-content-between mr-10 w-full self-center">
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
