import { useEffect, useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";
import { useParams } from "react-router-dom";
import PostSection from "../components/PostSection";
import useLogin, { modalStorage } from "../stores/useStore";
import {
  closeToast,
  closeToastWithoutFade,
  showToast,
  toastWithoutFade,
} from "../utils/toastUtils";
import Confimation from "../components/Confimation";

function SpecificPostView() {
  const { isLoggedIn } = useLogin();
  const { setModalType, modalState, setModalState } = modalStorage();
  const { postID } = useParams();

  const [post, setPost] = useState([]);
  const [sections, setSections] = useState([]);

  async function handleDelete() {
    setModalType("Delete");
    setModalState(true);
  }

  useEffect(() => {
    toastWithoutFade("", "Loading");
    fetch(`http://localhost:8080/post/${postID}`)
      .then((res) => res.json())
      .then((data) => {
        setPost(data);
        setSections(data.Sections);
      })
      .catch((error) => {
        closeToastWithoutFade();
        console.log(error);
        showToast("Failed to load post", "Warning");
      });
    closeToast();
  }, []);

  return (
    <>
      {modalState ? <Confimation /> : <></>}
      <div className="flex flex-row min-h-fit">
        <Navbar />
        <div className="flex flex-col w-full pb-10">
          <Logo />
          <div className="flex flex-col w-[1080px] gap-5">
            <div className="flex flex-row justify-between">
              <TitleNav string={post.Title} />
            </div>
            <div className="flex flex-col gap-2">
              {Array.isArray(sections) &&
                sections.map((data, index) => (
                  <div key={index} className="flex">
                    <PostSection sectionData={data} />
                  </div>
                ))}
            </div>
            {isLoggedIn ? (
              <button
                onClick={handleDelete}
                className="text-xl font-semibold bg-red-600 w-1/3 h-[60px] place-self-end rounded text-white hover:text-black hover:bg-white duration-200 "
              >
                Delete post
              </button>
            ) : (
              <></>
            )}
          </div>
        </div>
      </div>
    </>
  );
}

export default SpecificPostView;
