import PropTypes from "prop-types";
import { Link } from "react-router-dom";

function TitleNav({ isLink, link, string }) {
  return (
    <>
      {isLink ? (
        <Link to={link}>
          <p className="text-3xl font-semibold duration-200 hover:text-[#FFA360]">
            {string}
          </p>
        </Link>
      ) : (
        <p className="text-3xl font-semibold">{string}</p>
      )}
    </>
  );
}

TitleNav.propTypes = {
  isLink: PropTypes.bool,
  link: PropTypes.string,
  string: PropTypes.string,
};

export default TitleNav;
