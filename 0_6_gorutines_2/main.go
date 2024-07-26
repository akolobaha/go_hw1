package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const workersNumber = 10
const timeout = time.Millisecond * 1500
const basePath = "./images/"

func DownloadFile(url string, filepath string) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer response.Body.Close()

	// Проверяем код статуса ответа
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка при скачивании файла: %s", response.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	return nil
}

func GenerateRandomUID(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Worker(num int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		uid, _ := GenerateRandomUID(10)

		fmt.Println("worker", num, "started job", j)
		time.Sleep(timeout)

		err := DownloadFile(j, basePath+uid+".jpg")
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		results <- j
	}
}

func main() {
	jobsNum := len(imgUrls)
	jobs := make(chan string, jobsNum)
	results := make(chan string, jobsNum)

	for w := 1; w <= workersNumber; w++ {
		go Worker(w, jobs, results)
	}

	for _, j := range imgUrls {
		jobs <- j
	}
	close(jobs)

	for range jobsNum {
		<-results
	}
}

var imgUrls = []string{
	"https://avatars.mds.yandex.net/i?id=b0bff4a72cb54a09d3d1c243978290180103bd89dbc2b530-12764650-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=8ba258086e0a7fef926b2146d405d5fa-5232255-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b4fce25d18791ddf4034ce35127f3f78036f3f3e-12991056-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=0c85412d0e8f8373e8586f4b05ba4949c64e7984-10414173-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=bdef6eb7357fe1b3bed9301c8923a02f67b2a2a5-10754985-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=578bf0f92ff112f23fa78111dcbcf82afc9f2310c1464d9c-5234460-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=6ba5ff19056527e16230ac57e9fb54b5a462d3609c5c870e-12441694-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=aeea98de51f8f785477e625892ef39a53efc3b8a-11543356-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=2afdf5808da1212afc77c08c0b67c01b85cd5ce6-11471706-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=f61556c4874a978d21fc1b5b613a859c71b17e55-10109607-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=39249e23346711dede210e68e45ecb48bc208f0f-12422712-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=882d784eb58624378e35b82074bb9c4e409afc24410f74e8-5281109-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=f874f0a61cdbb17d0da8b43b5c7cb16274587baf-9228595-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=70b531de7b373498c2a24fa65585139b2824ff51-7760121-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=ffda63c88d5385f89cbcd9b04399dce46089bb28-9215233-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4ed8af4d6ec779986387614025ededfcc09ffc5f-9053276-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=ccd35f03d1727f4cd211d94ca04cddcb7f11c613-9246694-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4f4c3dd31a1cf408d72b34260faa3f6a-4907671-images-taas-consumers&n=11&ref=rq",
	"https://avatars.mds.yandex.net/i?id=b4736d109d018b8186e243a1e87d2d52-5235737-images-taas-consumers&n=11&ref=rq",
	"https://avatars.mds.yandex.net/i?id=c575af078acb181b18170e71df3f4b35-5277051-images-taas-consumers&n=11&ref=rq",
	"https://avatars.mds.yandex.net/i?id=72fcbf30b820d23797dedf8038201fe2-5032979-images-taas-consumers&n=11&ref=rq",
	"https://avatars.mds.yandex.net/i?id=22b13d20d1353651b27f556bdcc98b55-5220211-images-taas-consumers&n=11&ref=rq",
	"https://avatars.mds.yandex.net/i?id=cc9d4c6ad115f67e0aa62dd7c654ae35-5324191-images-taas-consumers&n=11&ref=rq",
	"https://avatars.mds.yandex.net/i?id=bf61b120081757d00d8e4f4fe58d63e4-5239606-images-taas-consumers&n=11&ref=rq",
	"https://avatars.mds.yandex.net/i?id=86e5c6ea589d0014d3db0055a18882caedcf4904-8243435-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=5bc729f30c151b5db56da3bb157ba1780433eb69-10851049-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=272c27bf61783a2b265a5968ca1b9a758d58b268-9683462-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=71b43855f97887cf1ffe41600267b606bd25a5f4-5234135-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=a206f5b8e84d185348290f2b025dd7c26368d4e5-8102231-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=f2d5d3b23ce424156d6dd144e7b8d85232074dd8-11541841-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=3a115c102553f384eb2e8ff1f989266660d8d59e-12423030-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=2babf49d66e1588e9fe3dd0e8107229181ce7c6f8babc048-4547856-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=db8a35e6a182304fbfb547116723d53360386034-5209934-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=158d3a08ffe9c7731eb0563438799bb80d9ac60d-12421110-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=143a5483a72e55c8f98a11e806cfdb629dd3db29-5115049-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=42efed9fa4cd59377f0a760baaea4a06d789b60c0ba784d8-12718238-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=3afa8ad32481579f4fca28aeb0433a10705e4ee0-12542812-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=f980c2d7c567622b5ccf90675ca14e0358870050-10255460-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=fa737a4e373f80a1926ff758c12de367-4570726-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4c2d15f6dec7a32d15dd86989bb29a19f00ad33e-6235060-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=634f7ce1ea56e043484ad3b445915bf816250708-10100414-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=1c358e9b9f85ce133cbb65ea1e4683cf9012a343-9287521-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=203a24410ebedf9826c474292733e599ae0f51eb-4120702-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=09f07c01334daeaff62dac616bd166473e595964ec4b74a3-12855379-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=c5491081814c0ab35740a91a04eeb667e554f6ebefe92146-13266420-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=2a00000179ed9d3eab12f524d1c30224a5c0-3940630-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=6785f35a42b05f25d86cff820446a7bc33859937-10995513-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=79c9b2ef32994ae1e9f3bbc5a43a254cc2e7ae14d51448eb-12346182-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=8ca6d255d09aefbb47ba2af7bea04c723e328b26-11918841-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=fccd6a383a477240ad5efb6ca92d89f78011a64f-6947113-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=93b99443c3a841d7890811146f07d7ec600b868d-12366556-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=0bf46ee7a2b65c140b73e3d92d5106eb73694f4d-9853586-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=5760bea0ee0ff996754a5c0da882343d1f5650ccc983a5c6-11922792-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=8dd926e4c9767c644386d1d0df20c416f64a4fe7-12810245-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4cbe3938476a447fe9710ea536dc90bd766a23a0-12373778-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=ecba443facdd5831afca7804a524509d5ec09618-9233745-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=19305632d6b9106167b15dc0bd25bfce33f08727-10995513-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=428ff01c375bf04553de35ac531df14ba4417950-12149948-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4fa377af3b536fbe499d3a72420afb1c94f3f6a3-12631601-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=29b65dfd4bd984b18bfd0348c4d40067e008b3c7-12925719-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=684ecd32ced6492edaaa7b89257f2bb24a19390c-12422360-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b5c0be0aa6afbd67c89f09e85b98468a98a95ee4d8146e81-12421956-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=bcd4972f42ad1c5508ebf8a224bf58d39d2efc6227c4f1af-11380463-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=79706fa1c61ae56bbfd357d13fcad30de9f9e7fb-7574298-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=5171f3d757e6e134216d0e3ebbc0a381cfe42173a449c866-5118451-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4d614c93b3be4e84859fcef64eb5e1ef090f3442-7452498-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=c7a77d5c68e780bd0a6c4eb66223d25a4a9ebda50b763d80-10681994-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b0cc0facbc9d8cd28b89c424e88bd78d494eb8cef4afb965-4570132-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=ea7a6e98c8af539139a671dc7aab51f1a2791a68-7543473-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=c6a9a4558c48bb6f52584c04d3043d9759a7b5b1-8209628-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=2925f953ef3deb34a6e594217199b976c4d704d6-10931829-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=a6c154ec370b318b10783ade107b772b8f670621-8210619-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=f3e76c7247707de61f2db33f585f568efcb244de-10093836-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=ac03a72d92b67d0a5e8ac2532a112b43d3a3ee89df2dfb77-5869613-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=de40aa33b6f7c3281ea36fe82ba62656a0b3a8cc-12914734-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=24a4454917d76b2a739d449c66dd8a89537fd5e2-11401793-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=e56e2c570352a6407d4988f17af7961aa2717dc2-4615722-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=5fb8484d86e48e9573ad58a1ddd17edfbbe836e327555462-5230955-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=e3c26a3ba06c33cf73ee1f7b58a44fe8697252ff-10085718-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=2ad375bc271c295e2f8ffb69738f8f8c4add2a88-12511705-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=d6b511e5a7248c6c8fd8937760057a4c58d6027f-10932765-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=6e695a16a4e731221713fa8e76246c8706645efd-10576628-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4bd90ef9122b7508c0053435667e84a768a851bc-7947996-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=a57d46d19ce3a80ec0cff40cf5d2cb2d3673a0035f5c9c21-11374844-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=6dd0bed677af3c2f65b4e302e582e4a07f5e104f-9067973-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=6096ffcc0eb41afaaa4154fbc4b34fd83af081c9-10959457-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=9ccab50b1d9d1520baf7db54685564eeaa631081-5236166-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=d1ae78733cb6c2f0f4debdd75866d84fc215b3f9-11950916-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=3a74f12d47837b19d8a55f4f63e8511dacdfe2c2-9866196-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=2616d21c92be7190b8e703c6144a819716a912f2-4601270-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=1f9519ea045a8def14efb17d0c394d73866ccade-2455092-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=caecc7f8403fe2f0d197f4dbdbae7a1d89f7e03d31f5e07d-12473832-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=99315d0c64cdacf51bb3e6a56706a2283ed25b91-12606451-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=4619a133f4b7b4c3a05fbe8a7344e1593ce21fa5-2361655-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=d86a92d7acca71475fa0cefd3d31f5796511e8d8-7544603-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b6443f030bb1527bd13653519dbe23fa3c6c021394c94a43-4262069-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=3e5bddbb6aeb2c96b5e7eafb5dd460c1139e552255fb5ead-5250028-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=7c026bd0af656b316013e3e6a0aa59fe48556a45-10896002-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=9d15b9fef6ba43ad03601446804b45df47fa05efe6c2175d-8497639-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=2fd11c41118f45037b7768b373d7b430f8af0257-10517487-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=5af9a29ccc91f5d8b9cc1abf6c95aeb2c872a1c1-10654381-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=8761fe837d1e0bdea280b5dac4861f3dd44888c1-5263021-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=95a91a830e1e2ae3bf2bc5f08ae170e1686f39e3-10354912-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=a00d0f182dfce3f2465168c4be0f3f5b48637db4-5293797-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=decac9274780adffe83c9d7d54474f8bf0ef2e0f-4769309-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=87766318d2f68fda37ea06e1548f7720ab798333-12845884-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=ecbe233b67156624a58381a3233dd04fcc7b946b-5336359-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=35b839f41f8df6acceefd8f1b8750973c48b8a1a-9720926-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=7a1eeb8b4d30dd3a5eeaf9cf59d500468906ea64-9866196-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=dfb841028c480ec913c2d28bb1d7be90277c24f8e9823784-4055830-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=23688bd6f8f1dceb54f2ca40d92126cf-5350095-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=7a57a4f048b64b2fec0493ac0b63f84f96a1e7d8-7761368-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=579d1b39c157206dc121c2f1219fc82b9405f55a-9106703-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=022d562d3d7b05237ae9fd1147b61b8923e15d25-12496607-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=e3c457539a8044bf7df5bbb03bd312d2ab37b2ae-5534042-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=d77e56b7be06ee03d85d659b08a9a4a4d7f1a94e-11907031-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=34b3d5a8d0f0adc21d01a0895706a5ed2b2d40da-8767819-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=32e9da1f81b846739bba3097b875f18f913dcd3b-12513999-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=ad7f627c850c40c88c537665e6dd55c2fa371a33-5151070-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b8eecae36b01b43373a573e506c431a6d2a97941-10371233-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=e7e6dee62462b0ba95833d4d57e7e8e5ecff199b-4835514-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b2b6b954e2902ce1eded2a319389803fa656a1a2-4473780-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=a3323b7826cde9e3ce1a74b1c124fd3089f07477eff574e3-5451798-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=7cfede6529b5307c6316e9e99d894b2dcdcc143e-8236365-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=db09772c36a60bd7632b933024e119d561dc02b8-5070325-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=826d7e1df730c87de62fc6c802f29dbf00b1e793-5334840-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=da8695a0161437339845fd8e026fafa94bbdafef-13604457-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=8f03626cbce40ad9d1fe954afdf5b673dccf9acc70dfe2b3-3719127-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=6cf02dd7cb4de159dd7fba9193ad06a2-5648144-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b8bfdf49a2a1f7fac62b77547302e599c48b531c-4824334-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=204f2fe58868a005668d6aac3482cf4ad6602524-12647631-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=a6e19fa719db7cacb65c81a6cec22d869e71808d-12609997-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=a8bbe6daa9d36d504a5d2471bde5dff9d69ae76a-12638015-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b6ee255fd3f5129f4b9e0eee52412e56dc795df7-12802892-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=598a854dd370ae184b909ab494bb45ff85c501f76f7baaa2-5303267-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=eeec4a0d7f988c61e23ed65175850051_sr-9868376-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=1a4d0c88081250161f46d1f6f002ebd3c14478dfad54cf9c-5013592-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=f2d423d9d56bdfcb83ba790122ad8bc4adc03b3cd27b1cf6-12384509-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=32073ee72a1a3edce4668256401b017c6b4cbb93bf024471-10263600-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=95b3c8becdeafc67e798e1b2c4931177fffab9f4-11395806-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=1c624e88c2cb125550011c709735c5c7253e9432-10190941-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b04d6f49fe7f30a441306bc49e533fd9f003b41a45f31f20-4290907-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=bdd6eaabfe45d402fd8d10c2ad26591e5fb6a1a33ac1d41d-9211188-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=34bf25deadafe2ce44393fd314e8325d09054880844c3101-12992500-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=936ef3028a9da8f74ec05ac8483f08de-4547856-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=b3b2d78a1d2c3cbc84277dbf070a6de25945d70b-4345430-images-thumbs&n=13",
	"https://avatars.mds.yandex.net/i?id=fb6ecc1e8d26f0cd84531aea2d6a764bf36241c1-10385090-images-thumbs&n=13",
}
