import TitleNav from "./atomic/TitleNav";
import useLogin, { modalStorage } from "../stores/useStore";
import Confirmation from "../components/Confimation";

function Navbar() {
  const isLoggedIn = useLogin((state) => state.isLoggedIn);
  const { modalState, setModalState, setModalType } = modalStorage();

  function handleLogout() {
    setModalType("Logout");
    setModalState(true);
  }

  return (
    <>
      {modalState ? <Confirmation /> : <></>}
      <div className="flex w-[390px] justify-end pt-40 pr-14 mr-20 border-r-4 border-[#FFA360] min-h-screen">
        <div className="flex flex-col gap-4 text-right place-items-end">
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
