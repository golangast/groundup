package kvalparse

//multiline unicode string
//http://www.madore.org/~david/misc/unitest/

var bigStringOne = `The following lines are the first chapter of the Qur'an (note that the text runs right to left, and should probably be aligned on the right margin):

بِسْمِ ٱللّٰهِ ٱلرَّحْمـَبنِ ٱلرَّحِيمِ

ٱلْحَمْدُ لِلّٰهِ رَبِّ ٱلْعَالَمِينَ

ٱلرَّحْمـَبنِ ٱلرَّحِيمِ

مَـالِكِ يَوْمِ ٱلدِّينِ

إِيَّاكَ نَعْبُدُ وَإِيَّاكَ نَسْتَعِينُ

ٱهْدِنَــــا ٱلصِّرَاطَ ٱلمُسْتَقِيمَ

صِرَاطَ ٱلَّذِينَ أَنعَمْتَ عَلَيهِمْ غَيرِ ٱلمَغضُوبِ عَلَيهِمْ وَلاَ ٱلضَّالِّينَ

Here is what the above might look like if your browser supports the Arabic block of Unicode:

[Seven verses in Arabic]

And here is a transcription of it:

bismi ăl-la'hi ăr-raḥma'ni ăr-raḥiymi

ăl-ḥamdu li-lla'hi rabbi ăl-"a'lamiyna

ăr-raḥma'ni ăr-raḥiymi

ma'liki yawmi ăd-diyni

'iyya'ka na"budu wa-'iyya'ka nasta"iynu

ĭhdina' ăṣ-ṣira'ṭa ăl-mustaqiyma

ṣira'ṭa ăllaḏiyna 'an"amta "alayhim ġayri ăl-maġḍuwbi "alayhim wala' ăḍ-ḍa'lliyna

A rough translation might be:

In the name of God, the beneficient, the merciful.

Praise be to God, lord of the worlds.

The beneficient, the merciful.

Master of the day of judgment.

Thee do we worship, and Thine aid we seek.

Lead us on the right path.

The path of those on whom Thou hast bestowed favors. Not of those who have earned Thy wrath, nor of those who go astray.`

var bigStringTwo = `The following are the two first lines of the Analects by Confucius:

子曰：「學而時習之，不亦說乎？有朋自遠方來，不亦樂乎？ 
人不知而不慍，不亦君子乎？」

有子曰：「其為人也孝弟，而好犯上者，鮮矣； 
不好犯上，而好作亂者，未之有也。君子務本，本立而道生。 
孝弟也者，其為仁之本與！」

Here is what the above might look like if your browser supports the CJK block of Unicode:

[Two lines in Chinese]

And here is the transcription of it:

Zǐ yuē: “Xué ér shī xí zhī, bú yì yuè hū? Yoǔ péng zì yǔan fānglái, bú yì lè hū? Rén bù zhī, ér bú yùn, bú yì jūnzǐ hū?”

Yóuzǐ yuē: “Qí wèi rén yě xiàodì, ér hàofànshàngzhě, xiān yǐ; bú hào fànshàng, ér hàozuòluànzhě, wèi zhī yóu yě. Jūnzǐ wù běn, běn lì ér dào shēng. Xiàodì yé zhě, qí wèi rén zhī bén yǔ!”

A rough translation might be:

The Master [Confucius] said: “To study and to practice, it is is a joy, isn't it? When friends come from afar, it is a pleasure, isn't it? If one remains unknown and isn't hurt, isn't one an honorable man?”

Master You said: “Few of the men who act well filially and fraternally are also fond of offending their superiors; men who are not fond of offending their superiors, but who like to cause trouble, such do not exist. The honorable man concerns himself with the foundations. Once the foundations are established, the Way is born. Is not acting well filially and fraternally the foundation of humanity?”`
