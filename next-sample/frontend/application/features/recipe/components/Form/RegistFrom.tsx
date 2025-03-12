"use client"; 

import {
  faFileLines,
  faCarrot
} from '@fortawesome/free-solid-svg-icons';


import React, { useState }  from 'react';
import { useRouter } from 'next/navigation'

import { APP_DATA } from '@/constants/appdata'

import TextForm from '@/components/elements/textForm/TextForm';
import TextFormArea from '@/components/elements/textFormArea/TextFormArea';
import Button from '@/components/elements/button/Button';

import { faPlus } from "@fortawesome/free-solid-svg-icons";
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

  const [postedData, setPostedData] = useState('')

  var tempMaterials: Material[];
  var uniqueItemsMap: Material[] = [];

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

  
  const postRecipe = async () => {
    // レシピ情報を送信するリクエスト
    const requestData: RecipeRequestData = {
      title: title,
      content: content,
      materials:materials,
    }
  
console.log(JSON.stringify(requestData));
    await fetch('http://localhost:9000/recipe/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestData),
    }).then(response => response.json())
    .then(data => {
        // .thenは成功した時の処理を示す場合に使う。
        console.log('Success:', data);
        alert('レシピの投稿に成功しました');
        router.push('/')
    })
    .catch((error) => {
        // .catchは失敗の時の処理を示す場合に使う。
        console.error('Error:', error);
    });


  }



  // この下からリターンの中身
    return (

        <div className="relative p-3">
 
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
              disabled={!title || !content || !materials}
              type="submit">
              登録
              
            </Button>
          </div>

        </div>

      
    );

};

export default RegistFrom;