import { ref, uploadBytesResumable } from "firebase/storage";
import { storage } from "../../firebase.config";

export async function uploadImg(img) {
  return new Promise((resolve, reject) => {
    const storageRef = ref(storage, "images/" + img.name);

    console.log(storageRef);

    const uploadTask = uploadBytesResumable(storageRef, img);

    uploadTask.on(
      "state_changed",
      (snapshot) => {
        const progress =
          (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
        console.log(progress + "% done");
      },
      (error) => {
        reject(error);
      },
      () => {
        resolve();
      }
    );
  });
}
