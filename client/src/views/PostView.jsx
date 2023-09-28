import { useEffect, useState } from "react";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";
import formatPublishedDate from "../utils/dateFormatter";
import { Link } from "react-router-dom";
import Logo from "../components/Logo";

function PostView() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/posts")
      .then((response) => response.json())
      .then((data) => {
        setPosts(data);
      })
      .catch((error) => {
        console.error(error);
      });
  }, []);

  return (
    <>
      <Logo />
      <div className="flex flex-row h-full w-full">
        <Navbar />
        <div className="flex flex-col h-3/4">
          <TitleNav string="Entries" />
          <div className="flex gap-5">
            <div className="flex flex-col gap-4">
              {Array.isArray(posts) &&
                posts.map((post, index) => (
                  <Link to={`/posts/${post.ID}`} key={index}>
                    <p className="text-xl duration-200 hover:text-[#FFA360]">
                      {post.Title}
                    </p>
                  </Link>
                ))}
            </div>
            <div className="flex flex-col gap-4">
              {Array.isArray(posts) &&
                posts.map((post, index) => (
                  <p key={index} className="text-xl">
                    {formatPublishedDate(post.Published)}
                  </p>
                ))}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default PostView;
