import TitleNav from "./atomic/TitleNav";

function Navbar() {
  return (
    <>
      <div className="flex w-[320px] h-[600px] content-end justify-end p-10">
        <div className="flex flex-col gap-4 text-right h-5/6">
          <TitleNav isLink={true} link="/" string="Home" />
          <TitleNav isLink={true} link="/portfolio" string="Portfolio" />
          <TitleNav isLink={true} link="/posts" string="Posts" />
          <TitleNav isLink={true} link="/contact" string="Contact" />
        </div>
      </div>
    </>
  );
}

export default Navbar;
