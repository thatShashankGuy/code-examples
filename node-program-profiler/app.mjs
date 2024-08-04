import fs from 'fs'
import sharp from 'sharp';
import archiver from 'archiver';



async function compAndZip(){
    try {
        await sharp('image.jpg').resize(800,600,{fit:'inside'}).toFile('compressed_image.jpg')

        const output = fs.createWriteStream('zipped_image.zip')
        const archive = archiver('zip',{zlib:{level:9}})

        output.on('close',()=>{
            console.log(archive.pointer()+' total byte')
            console.log("zip succesful")
        })
    } catch (error) {
        console.log(`Error while archiving and compressing : ${error.message}`)
    }
}

compAndZip()

