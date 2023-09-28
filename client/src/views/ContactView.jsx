import Logo from "../components/Logo";
import LinkedInIcon from "../assets/icons/LinkdIn.png";
import EmailIcon from "../assets/icons/Email.png";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";

function ContactView() {
  return (
    <>
      <Logo />
      <div className="flex flex-row h-full w-full">
        <Navbar />
        <div>
          <TitleNav string="Contact" />
          <div className="flex flex-row gap-3 pt-4">
            <div className="flex flex-col gap-5">
              <a href="https://www.linkedin.com/in/rizky-soeharto-aa33b0222/">
                <img
                  src={LinkedInIcon}
                  alt="LinkedIn Icon"
                  className="w-[32px] h-[32px]"
                />
              </a>
              <a href="mailto:soehartorizky@gmail.com">
                <img
                  src={EmailIcon}
                  alt="Mail Icon"
                  className="w-[39px] h-[32px]"
                />
              </a>
            </div>
            <div className="flex flex-col gap-5">
              <a
                className="text-2xl font-semibold ml-3"
                href="https://www.linkedin.com/in/rizky-soeharto-aa33b0222/"
              >
                LinkedIn
              </a>
              <a
                className="text-2xl font-semibold ml-3"
                href="mailto:soehartorizky@gmail.com"
              >
                soehartorizky@gmail.com
              </a>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default ContactView;
