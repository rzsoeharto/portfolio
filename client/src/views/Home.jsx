import { Link } from "react-router-dom";
import ProfilePhoto from "../assets/ProfilePhoto.png";
import LogoImage from "../assets/Logo.png";

function HomeView() {
  return (
    <>
      <div className="flex flex-row w-full h-full place-content-center items-center gap-48">
        <div className="place-items-center">
          <img src={ProfilePhoto} alt="Photo of Rizky" />
          <img src={LogoImage} alt="" className="self-center" />
        </div>
        <div className="flex flex-col gap-10 text-3xl font-semibold">
          <div className="flex flex-row gap-2">
            <p> Hi! My name is</p>
            <p className="text-[#FFA360]">Rizky</p>
          </div>
          <div className="w-[550px]">
            <p>
              I am an aspiring Software Engineer, focused in Back end
              Development. And I try to <Link to="/posts">write </Link>
              too (sometimes).
            </p>
          </div>
        </div>
      </div>
    </>
  );
}

export default HomeView;
