"use client"; 

import {
  faFileLines,
  faCarrot
} from '@fortawesome/free-solid-svg-icons';


import React, { useState, useRef }  from 'react';
import { useRouter } from 'next/navigation'

import { APP_DATA } from '@/constants/appdata'

import TextForm from '@/components/elements/textForm/TextForm';
import TextFormArea from '@/components/elements/textFormArea/TextFormArea';
import Button from '@/components/elements/button/Button';

import { faPlus,faCamera } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { useInputValue, useInputContent } from '@/hooks/useInputValue';


type Material = string

// レシピ投稿時に送信するデータの型
type RecipeRequestData = {
    title: string;
    content: string;
  //  image: string;
    materials?: Material[];
};



const RegistFrom = () => {

  const router = useRouter()

  const [title, updateTitle] = useInputValue('');

  const [content, updateContent] = useInputContent('');
 //const [material, updateMaterial] = useInputValue('');

  // フロントで一時的にタグを保持するためのstate
  const [tempMaterial, updateTempMaterial, setTempMaterial] = useInputValue('');

  // 送信するためのタグ配列を保持するためのstate
  const [materials, setMaterials] = useState<Material[]>([]); 


  const imageForm = useRef<HTMLInputElement>(null);
  const [image, setImage] = useState<File | null>(null);
  const [preview, setPreview] = useState<string | null>(null); //画像のプレビューを表示するためのstate
  const [errMessage, setErrMessage] = useState('')

  var tempMaterials: Material[];


  // 材料追加のロジックをまとめた関数
  const addMaterial = () => {


    // 材料の数が APP_DATA.MATERIAL_MAX_NUM 以下の場合のみ、タグを追加できるようにする。
    if (tempMaterial && materials.length <= APP_DATA.MATERIAL_MAX_NUM) {
      console.log("addMaterial");
      console.log(materials);
      console.log([...materials, tempMaterial]);
      tempMaterials = [...materials, tempMaterial].filter((item, index, self) => 
        index === self.findIndex((t) => t === item)
     );


     // uniqueItemsMap = new Map(tempMaterials.map(item => [item, item])); 
     console.log("tempMaterials");
      console.log(tempMaterials);
      setMaterials(tempMaterials);
      setTempMaterial('');
    }
  };

  // 画像の形式バリデーションとプレビュー表示を行う関数
  const imageCheck = (fileList: FileList) => {
    if (fileList[0]) {
      const imageTypes = ['image/jpeg', 'image/png', 'image/gif'];
      if (!imageTypes.includes(fileList[0].type)) {
        alert('許可されていないファイルタイプです。');
        return;
      }
      console.log(fileList[0]);
      setImage(fileList[0]);
      setPreview(URL.createObjectURL(fileList[0]));
    } else {
      setImage(null);
      setPreview(null);
    }
  };
  

  const createFormData = () => {       
    const formData = new FormData()
    if (image) {
      formData.append('file', image) 
    }
    formData.append('title', title)
    formData.append('content', content) // ポイント1！
    
    if (materials.length > 0) {
      formData.append("materials", JSON.stringify(materials));
    } 

    return formData
  }

  const postRecipe = async () => {


/*


   const url = 'http://localhost:9000/recipe/register'
    const data = await createFormData()   //formdataが作成されるのを待つ
    const config = {
      headers: {
        'content-type': 'multipart/form-data'
      }
    }
    axios.post(url, data, config)
    .then(response => {
      console.log('Success:', data);
        alert('レシピの投稿に成功しました');
       // router.push('/')
    }).catch(error => {
      console.log(error)
      alert('error');
    })
*/

    const data = createFormData()   //formdataが作成されるのを待つ  
    await fetch('http://localhost:9000/recipe/register', {
      method: 'POST',
     // headers: {
      //  'Content-Type': 'application/json'
     //   'content-type': 'multipart/form-data'
     // },
      body: data,
    }).then(response => response.json())
    .then(data => {

      if (data.error != undefined) {
        setErrMessage(data.error);
      } else {
        // .thenは成功した時の処理を示す場合に使う。
        console.log('Success:', data);
        alert('レシピの投稿に成功しました');
        router.push('/')
      }
    })
    .catch((error) => {
        // .catchは失敗の時の処理を示す場合に使う。
        console.error('Error:', error);
    });


  }
  



  // この下からリターンの中身
    return (

        <div className="relative p-3">

          <div className="mx-auto flex px-3 py-3 items-center justify-center  text-red-500">
              {errMessage}
          </div>

          <div className="relative">
            
          <input
            ref={imageForm}
            type="file"
            className="hidden w-full"
            accept="image/*" //画像のみを選択できるようにする

            onChange={(e) => {
              const fileList = e.target.files;
              if (fileList) {
                imageCheck(fileList);
              }
            }}
          />
            
            {preview ? (
            // 画像が選択されている場合は、プレビューを表示する
            <img
              className="mx-auto h-52 w-72 cursor-pointer rounded-2xl border-4 border-solid border-[#f4cdd8] object-cover shadow-md"
              onClick={() => {
                if (imageForm.current) {
                  imageForm.current.click();
                }
              }}
              src={preview}
              alt="プレビュー"
              width={300}
              height={200}
            />
          ) : (
            <div
              className="mx-auto flex h-52 w-72 cursor-pointer flex-col items-center justify-center rounded-2xl border-4 border-dashed border-gray-400 text-gray-500 hover:border-pink-500"
              onClick={() => {
                if (imageForm.current) {
                  imageForm.current.click();
                }
              }}
            >
              <FontAwesomeIcon icon={faCamera} className="text-8xl" />
              <div className="text-2xl">写真選択</div>
            </div>
            )}
          </div>
          <TextForm
            label="レシピタイトル"
            labelIcon={faFileLines}
            placeholder={"※ 必須(" + APP_DATA.RECIPE_TITLE_MAX_LENGTH + "文字以内)"}
            maxLength={APP_DATA.RECIPE_TITLE_MAX_LENGTH}
            value={title}
            onChange={updateTitle}
          
          />

          <div className="flex justify-center items-center">
            <TextForm
              label="材料"
              labelIcon={faCarrot}
              placeholder={"※ " + APP_DATA.MATERIAL_MAX_NUM + "個以内(" + APP_DATA.MATERIAL_MAX_LENGTH + "文字以内)"}
              maxLength={APP_DATA.MATERIAL_MAX_LENGTH}
              value={tempMaterial}
              disabled={materials.length >= APP_DATA.MATERIAL_MAX_NUM}
              onChange={updateTempMaterial}
            />
            <div className="ml-1 mr-4 mt-16">
              <Button 
                className="rounded-full bg-[#94C8AD] px-3 py-3 text-center text-sm text-white shadow-md transition duration-75 ease-in-out hover:bg-[#68B68D] active:scale-105 disabled:cursor-not-allowed disabled:bg-gray-300 disabled:hover:bg-gray-400"
                intent="secondary" 
                onClick={addMaterial} disabled={materials.length >= APP_DATA.MATERIAL_MAX_NUM}
                type="submit">
                <FontAwesomeIcon icon={faPlus} className="text-lg" />
              </Button>
            </div>
          </div>
          <div className="flex flex-wrap">
            {materials.map((material, index) => (
              <span key={index} className="mr-2"> 
                {material}
              </span>
            ))}
          </div>
          <TextFormArea
            label="作り方"
            labelIcon={faFileLines}
            placeholder={"※ 必須(" + APP_DATA.CONTENT_MAX_LENGTH + "文字以内)"}
            maxLength={APP_DATA.CONTENT_MAX_LENGTH}
            onChange={updateContent}
          
          />
          <div className="flex justify-center">
            <Button
              onClick={() => {
                postRecipe();
              }} 
              intent="primary" 
              disabled={!title || !content || materials.length <= 0}
              type="submit">
              登録
              
            </Button>
          </div>

        </div>

      
    );

};

export default RegistFrom;