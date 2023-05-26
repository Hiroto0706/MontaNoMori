<template>
  <v-container>
    <div class="content pa-10 my-10">
      <h1 class="content__title">ご利用規約</h1>
      <br />
      <p>
        もんたの森の素材をご利用頂く前に、利用規約をご一読ください！<br />

        <br />

        素材のダウンロードをもって利用規約に同意したものとみなします（規約同意ボタン等はございません）。
        <br />

        <br />

        <span class="font-weight-bold red--text"
          >個人利用の方：基本的に何してもOK！</span
        ><br />
        TwitteやLINEなどのアイコン、スマホの壁紙、ブログやホームページのアクセント、子供のお絵描き遊び、などなど！<br />

        <br />
        <span class="font-weight-bold red--text"
          >商用利用の方：どこにでも、何点でもご利用可能！</span
        ><br />
        商用利用や素材を加工しての利用はOK！
        広告やポップ、プレゼン資料でたびたびご利用頂けてます。
      </p>

      <br />
      <h2>これはOKです</h2>
      <br />
      <p>もうなんでもOK</p>

      <br />
      <h2>これはNGです</h2>
      <br />
      <p>とくにない</p>

      <br />
      <h1 class="content__title">よくある質問</h1>
      <br />
      <p>なんでもOKだから質問なんて来ません。</p>

      <br />
      <h1 class="content__title">このサイトについて</h1>
      <br />
      <p>
        運営者が絵を描くことが好きなので、その絵を世界にばら撒けたらいいなって思って作ったテキトーなサイトです。<br />
        描いた絵を投稿していきますので、是非使ってください。あなたの人生は360度変わります。
        <br />
        <br />
        このサイトに出会ったあなたはラッキーです。今後会うことはないでしょう。
        <br />
      </p>
      <div class="d-flex flex-column justify-end align-end">
        <v-img
          src="@/assets/about-me.png"
          lazy-src="@/assets/about-me.png"
          max-height="300"
          max-width="300"
          aspect-ratio="1"
        ></v-img>
      </div>
    </div>
  </v-container>
</template>

<style scoped>
.content {
  width: 100%;
  background-color: white;
  border-radius: 10px;
}

.content__title {
  display: block;
  width: 100%;
  background: linear-gradient(transparent 92%, #000000 92%);
}
</style>
