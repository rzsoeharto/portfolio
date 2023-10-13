import TitleNav from "./atomic/TitleNav";
import useLogin from "../stores/useStore";
import Cookies from "js-cookie";
import { useNavigate } from "react-router-dom";
import {
  closeToast,
  closeToastWithoutFade,
  showToast,
  toastWithoutFade,
} from "../utils/toastUtils";

function Navbar() {
  const isLoggedIn = useLogin((state) => state.isLoggedIn);
  const navigate = useNavigate();

  function handleLogout() {
    toastWithoutFade("Logging out", "Loading");
    const refresh = Cookies.get("Ref");
    fetch("http://localhost:8080/logout", {
      method: "POST",
      headers: {
        Authorization: refresh,
      },
    }).then((res) => {
      if (!res.ok) {
        closeToastWithoutFade();
        showToast("Failed to logout", "Warning");
        return;
      }
      closeToast(1000);
      navigate("/");
    });
  }

  return (
    <>
      <div className="flex w-[320px] h-[600px] content-end justify-end mr-32">
        <div className="flex flex-col gap-4 text-right h-5/6">
          <TitleNav isLink={true} link="/" string="Home" />
          <TitleNav isLink={true} link="/portfolio" string="Portfolio" />
          <TitleNav isLink={true} link="/posts" string="Posts" />
          <TitleNav isLink={true} link="/contact" string="Contact" />
          {isLoggedIn ? (
            <button
              onClick={handleLogout}
              className="bg-transparent text-3xl font-semibold hover:text-[#FFA360] duration-200"
            >
              {" "}
              Logout
            </button>
          ) : (
            <></>
          )}
        </div>
      </div>
    </>
  );
}

export default Navbar;
