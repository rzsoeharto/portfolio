import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import Confirmation from "../components/Confimation";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import SectionSelection from "../components/SectionSelected";
import AddSection from "../components/AddSection";
import useLogin, { modalStorage } from "../stores/useStore";
import { closeToast, showToast, toastWithoutFade } from "../utils/toastUtils";
import { uploadImg } from "../utils/firebaseUpload";

function NewPostView() {
  const navigate = useNavigate();
  const { isLoggedIn } = useLogin();
  const [titleData, setTitleData] = useState("");
  const [sectionsData, setSectionsData] = useState([]);
  const [imageArray, setImageArray] = useState([]);
  const [sectionModal, setSectionModal] = useState(false);
  const { modalState, setModalState, setModalType } = modalStorage();

  function CancelPost() {
    if (titleData == "" && sectionsData == 0) {
      navigate("/posts");
      return;
    }
    setModalType("");
    setModalState(true);
  }

  function onChange(e) {
    setTitleData(e.target.value);
  }

  function handleDelete(id) {
    const data = sectionsData[id];
    if (data.SectionType == "Image") {
      setImageArray((state) => {
        const updatedImgArray = [...state];
        updatedImgArray.splice(id, 1);
        return updatedImgArray;
      });
    }
    setSectionsData((state) => {
      const updatedSection = [...state];
      updatedSection.splice(id, 1);
      return updatedSection;
    });
  }

  async function submitPost() {
    if (!isLoggedIn) {
      showToast("You are not logged in", "Warning");
      return;
    }

    if (titleData === "") {
      showToast("Title cannot be empty", "Warning");
      return;
    } else if (sectionsData.length === 0) {
      showToast("Content cannot be empty", "Warning");
      return;
    }

    try {
      const uploadPromises = imageArray.map((image) => uploadImg(image));
      await Promise.all(uploadPromises);
      showToast("Images uploaded", "Success");
    } catch (error) {
      showToast("Failed to upload images", "Warning");
      console.error("RETURNED ERROR:                    ", error);
      return;
    }

    toastWithoutFade("Saving", "Loading");

    const response = await fetch("http://localhost:8080/create-post", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Title: titleData,
        Sections: sectionsData,
      }),
    });

    if (!response.ok) {
      showToast("Something went wrong in the backend", "Warning");
      return;
    }

    closeToast(1000);
    navigate("/posts");
  }

  return (
    <>
      {modalState ? <Confirmation /> : <></>}
      <div className="flex flex-row h-screen">
        <Navbar />
        <div className="flex flex-col w-full">
          <Logo />
          <div className="flex flex-col w-full gap-3 pb-10">
            <div className="flex flex-row gap-5">
              <input
                id="Title"
                onChange={onChange}
                type="text"
                placeholder="Title"
                className="text-3xl text-[#303030] w-[740px] font-semibold focus:outline-none"
              />
              <button
                onClick={CancelPost}
                className="text-2xl duration-200 hover:text-[#FFA360] bg-transparent"
              >
                Cancel
              </button>
            </div>
            <div className="flex flex-col w-[740px] gap-2">
              <div className="flex flex-col gap-3">
                {Array.isArray(sectionsData) &&
                  sectionsData.map((section, index) => (
                    <div key={index} className="text-lg flex flex-row">
                      <SectionSelection
                        index={index}
                        sectionSelection={section}
                        setSectionsData={setSectionsData}
                        setImageArray={setImageArray}
                      />
                      <button
                        onClick={() => {
                          handleDelete(index);
                        }}
                        className="relative right-5 top-7 font-bold bg-transparent h-10 w-14"
                      >
                        X
                      </button>
                    </div>
                  ))}
              </div>
              <div className="flex flex-row gap">
                {sectionModal ? (
                  <AddSection
                    setSectionModal={setSectionModal}
                    setSectionsData={setSectionsData}
                  />
                ) : (
                  <button
                    onClick={() => {
                      setSectionModal(true);
                    }}
                    className="bg-[#d9d9d9] h-[60px] text-xl w-full rounded hover:bg-white duration-200"
                  >
                    + Add Section
                  </button>
                )}
              </div>
              <button
                className="h-[60px] bg-[#FFA360] rounded text-xl font-semibold hover:bg-white duration-200"
                onClick={submitPost}
              >
                Submit
              </button>
              <button
                className=""
                onClick={() =>
                  // console.log(
                  //   "Sections: ",
                  //   sectionsData,
                  //   "Images: ",
                  //   imageArray
                  // )
                  console.log()
                }
              >
                log the thing
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default NewPostView;
