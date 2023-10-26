import { toastStorage } from "../stores/useStore";
import Loader from "./atomic/Loader";

function Toast() {
  const { toastType, toastMessage } = toastStorage();

  let type;

  switch (toastType) {
    case "Loading":
      type = (
        <div className="flex flex-row bg-white drop-shadow-md gap-5 w-80 h-20 place-items-center absolute bottom-[2%] left-[0%] px-5 font-semibold">
          <Loader />
          <p className="text-lg">{toastMessage}</p>
        </div>
      );
      break;

    case "Warning":
      type = (
        <>
          <div className="flex flex-row drop-shadow-md bg-white gap-5 w-80 h-20 place-items-center absolute bottom-[2%] left-[0%] px-5 text-red-600 font-semibold">
            <p className="text-lg">{toastMessage}</p>
          </div>
        </>
      );
      break;

    case "Success":
      type = (
        <>
          <div className="flex flex-row drop-shadow-md bg-white gap-5 w-80 h-20 place-items-center absolute bottom-[2%] left-[0%] px-5 text-green-600 font-semibold">
            <p className="text-lg">{toastMessage}</p>
          </div>
        </>
      );
      break;

    default:
      type = (
        <div className="flex flex-row bg-white drop-shadow-md gap-5 w-80 h-20 place-items-center absolute bottom-[2%] left-[0%] px-5 font-semibold">
          <p className="text-lg">{toastMessage}</p>
        </div>
      );
      break;
  }
  return (
    <>
      <div id="toast" className="hidden">
        {type}
      </div>
    </>
  );
}

export default Toast;
