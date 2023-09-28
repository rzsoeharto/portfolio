import { useEffect, useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";
import { Link, useParams } from "react-router-dom";
import PostSection from "../components/PostSection";

function SpecificPostView() {
  const { postID } = useParams();
  const [post, setPost] = useState([]);
  const [sections, setSections] = useState([]);

  useEffect(() => {
    fetch(`http://localhost:8080/post/${postID}`)
      .then((response) => response.json())
      .then((data) => {
        setPost(data);
        setSections(data.Sections);
      })
      .catch((error) => {
        console.error(error);
      });
  }, []);

  return (
    <>
      <Logo />
      <div className="flex flex-row h-min">
        <Navbar />
        <div className="flex flex-col w-[1080px]">
          <div className="flex flex-row justify-between h-min">
            <TitleNav string={post.Title} />
            <Link to="/posts/">
              <p className="text-2xl duration-200 hover:text-[#FFA360]">
                Close
              </p>
            </Link>
          </div>
          <div className="flex flex-col gap-2">
            {Array.isArray(sections) &&
              sections.map((data, index) => (
                <div key={index} className="flex">
                  <PostSection sectionData={data} />
                </div>
              ))}
          </div>
        </div>
      </div>
    </>
  );
}

export default SpecificPostView;
