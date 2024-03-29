import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";

function PortfolioView() {
  return (
    <>
      <div className="flex flex-row h-screen">
        <Navbar />
        <div className="flex flex-col w-full">
          <Logo />
          <div className="flex flex-col w-[1000px]">
            <TitleNav string="Portfolio" />
            <div className="w-4/5">
              <p className="text-2xl">
                As of right now. This is website is the only thing I am proud to
                share. However, you can check out my{" "}
                <a
                  href="https://github.com/rzsoeharto"
                  className="text-[#FFA360]"
                >
                  Github{" "}
                </a>
                to see my past projects.
              </p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default PortfolioView;
