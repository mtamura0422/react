
import React from 'react';
import * as Recipe from "@/features/recipe/components/Index"

import {
  faHeart as solidHeart,
  faClockRotateLeft,
  faCircleUser,
  faUtensils,
  faCarrot,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
export type Material = {
  id: number
  name: string
}


async function getData(id: number) {

  // var res = await fetch("http://localhost:9000/recipe/list")
  var res = await fetch("http://backend-service:9000/recipe/" + id)
 
  return res.json();
 }
 



export default async function Detail({
    params,
  }: {
    params: { id: number }
  }) {

    const resolvedParams = await params; // ✅ `params` を `await` してから使う
    const id = resolvedParams.id;



    const recipe: Recipe.RecipeData = await getData(id);
    
    var image: string;

    if (recipe.image == undefined) {
      image = "/images/noimage.png"
    } else {
      image = recipe.image;
    }
    

    return (
    <div className="select-none">
      <div className="mx-4 mb-5 mt-2 inline-block whitespace-pre-wrap break-all rounded-md text-2xl font-bold text-pink-500 lg:mx-7">
        {recipe.title}
      </div>

      <div className="relative mx-auto h-52 w-72 ">
            <img
              src={image}
              alt="レシピ画像"
              className="relative h-full w-full rounded-2xl border-4 border-[#f0bee2] bg-white object-cover shadow-md"
            />
      </div>

      <div className="mb-2 mt-10 text-lg font-semibold text-pink-500">
        <FontAwesomeIcon icon={faUtensils} className="mr-2 w-5" />
        作り方
      </div>
      <div className="rounded-md bg-[#f9f2e8] px-2 py-2 shadow-md">
        {/* 改行や空白を正しく表示させる処理 */}
        <div className="whitespace-pre-wrap break-all">{recipe.content}</div>
      </div>
      <div className="mb-2 mt-10 text-lg font-semibold text-pink-500">
        <FontAwesomeIcon icon={faCarrot} className="mr-2 w-5" />
        材料
      </div>
      {/* レシピに紐づくタグを表示 */}
      <div className="flex flex-wrap">
        {recipe.materials.map((material) =>
          <span className="mr-2">{material}</span>
        )}
      </div>
      {/* 作成ユーザーを表示 */}
      <div className="mt-10 text-center text-sm text-gray-500">
        <FontAwesomeIcon icon={faCircleUser} className="mr-2" />
        作成者：
        <span
          className="inline-block cursor-pointer rounded-md px-0.5 py-0.5 transition duration-75 ease-in-out hover:bg-[#FCCFA5] hover:text-orange-500 hover:shadow-md active:scale-105"
          
        >
          ユーザー名
        </span>
      </div>
      {/* 作成日と更新日を表示 */}
      <div className="mt-4 text-center text-xs text-gray-500">
        <FontAwesomeIcon icon={faClockRotateLeft} className="mr-2" />
        作成日： {recipe.updated_at}
      
      </div>
      {/* 現在ログインしているユーザーがレシピを作成したユーザーである場合に、編集ボタンが表示される。 */}
      
      
    </div>
      
    );
};
