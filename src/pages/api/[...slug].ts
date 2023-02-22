// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

function dogApiProxy(req :string) : Promise<JSON>{
  return fetch(`https://dog.ceo${req}`).then( res => res.json())
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<JSON>
) {
  let url = req.url
  if(url){
    let dog = await dogApiProxy(url)
    res.status(200).json(dog)
  }
  else {
    res.status(400)
  }
}
