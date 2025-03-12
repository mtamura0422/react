const recipes = [
  {
    id: 1,
    title: "レシピ1",
    img: "https://konbini-recipe.s3.ap-northeast-1.amazonaws.com/1711380068359-1685537767263-IMG_3907.jpeg",
    material: [
      {
        id: 1,
        name: "しらす"
      }
    ]
  },
  {
    id: 2,
    title: "レシピ2",
    img: "https://konbini-recipe.s3.ap-northeast-1.amazonaws.com/1711380068359-1685537767263-IMG_3907.jpeg",
    material: [
      {
        id: 2,
        name: "かぼちゃ"
      },
      {
        id: 3,
        name: "白菜"
      }
    ]
  
  }
]


export async function GET(request: Request) {
  return Response.json(recipes, { status: 200 })
}