import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import Confirmation from "../components/Confimation";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import SectionSelection from "../components/SectionSelected";
import AddSection from "../components/AddSection";
import { modalStorage, postStorage } from "../stores/useStore";
import Cookies from "js-cookie";
import Toast from "../components/toast";
import { showToast, toastWithoutFade } from "../utils/toastUtils";

function NewPostView() {
  const navigate = useNavigate();
  const { titleData, setTitle, sectionsData, removeSection } = postStorage();
  const { modalState, setModalState, setModalType } = modalStorage();
  const [sectionModal, setSectionModal] = useState(false);

  function CancelPost() {
    if (titleData == "" && sectionsData == 0) {
      navigate("/posts");
      return;
    }
    setModalType("");
    setModalState(true);
  }

  function onChange(e) {
    setTitle(e.target.value);
  }

  function handleDelete(index) {
    removeSection(index);
  }

  function submitPost() {
    const access = Cookies.get("Auth");

    if (!access) {
      showToast("You are not logged in", "Warning");
      return;
    }

    if (titleData == "") {
      showToast("Title can not be empty", "Warning");
      return;
    } else if (sectionsData == 0) {
      showToast("Content cannot be empty", "Warning");
      return;
    }

    let id = toastWithoutFade("Saving", "Loading");

    fetch("http://localhost:8080/create-post", {
      method: "POST",
      body: JSON.stringify({
        Title: titleData,
        Sections: sectionsData,
      }),
      headers: {
        "Content-Type": "application/json",
        Authorization: access,
      },
    })
      .then((res) => {
        if (!res.ok) {
          showToast("Something went wrong in the backend", "Warning");
          return;
        }
        id.className = id.className.replace("block ", "hidden show");
      })
      .catch(() => {
        showToast("Unable to connect to server", "Warning");
        return;
      });
  }

  return (
    <>
      {modalState ? <Confirmation /> : <></>}
      <Logo />
      <div className="flex flex-row">
        <Toast />
        <Navbar />
        <div className="flex flex-col w-[895.5px] gap-3 pb-10">
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
                    />
                    <button
                      onClick={() => {
                        handleDelete(index);
                      }}
                      className="font-bold bg-transparent h-10 w-14"
                    >
                      X
                    </button>
                  </div>
                ))}
            </div>
            <div className="flex flex-row gap">
              {sectionModal ? (
                <AddSection setSectionModal={setSectionModal} />
              ) : (
                <button
                  onClick={() => {
                    setSectionModal(true);
                  }}
                  className="bg-[#d9d9d9] h-16 text-xl w-full rounded hover:bg-white"
                >
                  + Add Section
                </button>
              )}
            </div>
            <button
              className="h-16 bg-[#FFA360] text-xl font-semibold hover:bg-white"
              onClick={submitPost}
            >
              Submit
            </button>
            <button className="" onClick={() => console.log(sectionsData)}>
              log the thing
            </button>
          </div>
        </div>
      </div>
    </>
  );
}

export default NewPostView;
