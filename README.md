# GetGhibliImages  

### スタジオジブリの公開画像をダウンロードするプログラムです  
Go言語で記載されています。
exeファイルもUPしてあるので実行するだけで画像をダウンロードできます。

---

### 使い方（実行ファイルからすぐにダウンロードしたい人）
exeファイル
特定の映画タイトルの公開画像(50枚)を取得します。
exeをダウンロードしてコンソールなどから実行してください。
実行するとexeファイルと同じディレクトリに画像ファイル（例：chihiro001.jpg〜chihiro050.jpg）が保存されます  

---
### exeファイルと映画タイトル  

| ファイル名 | 映画タイトル |  
|:-----------:|:------------:|
| GetGhibliImage_GedSenki.exe    | ゲド戦記 |  
| GetGhibliImage_Kaguyahime.exe  | かぐや姫の物語 |  
| GetGhibliImage_Karigurashi.exe | 借りぐらしのアリエッティ |   
| GetGhibliImage_Kazetachinu.exe | 風立ちぬ |  
| GetGhibliImage_Kokurikozaka    | コクリコ坂から |  
| GetGhibliImage_Marnie.exe      | 思い出のマーニー |  
| GetGhibliImage_Ponyo.exe       | 崖の上のポニョ |  
| GetGhibliImage_SentoChihiro.exe  | 千と千尋の神隠し |  

---
### ソースコード  
GetGhibliImage.go
映画ごとにURLが異なるのでmovieTitleを修正してビルドするだけでよいです。  
ビルド方法「go build GetGhibliImage.go」

