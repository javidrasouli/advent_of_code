package main

const Buffer = "hlfhfzffqnnrlnnvnmmgbgwgttbppcrcnnmdmfdmmgwwrrqnrrscrctcbttvcvtvvhchjhccjgjttmddplplqplqlbqlblrrbrvvprpff" +
	"pmmzpmpcczjzzbwwfssvrrvggncgncgcwczzswwqqjjflffpwfpwpbwpwpdpbpvvqffcfcjffjllncczfzzmhzzmddgdrgrwwjzzdjjsnjsjfsj" +
	"sjhhcchlccchqchhzzpnngdgndnpnppsdsggbvgvgpprqrqmmlzmzllvrrcvclcwczcqqcdcfcqqmmzbzdzdjzdjdmjdjzdjjcvjvcjvcvsslts" +
	"tttfbtftrfrlrdllrqqfssslccjdcjdjfdjfjqjnqjnjnrnddtnndtnnztzqztqztqzzpmzmggzrgrwwdqwdwcdwdnnmlmgmtmtstwssbffcncl" +
	"clnclcjjcjpcpqcpqcpqpmqqfccpcjppnspsnnzggnpntndtdqtthwhnhwnwllzhlzhlzzghzghhlhvhwhjhfjjcnjnvjnjvvqccdmmgddllnmn" +
	"rrdtdnncggfhgfglfgfmfnnpvvggznnwvnwwfgghrrfwrwzwszzzldzdldhlhblhblhbbbgjgsjggmqqmrrzggrhhwpwdpwdpwplpgpbggtssqf" +
	"fbqfbqbnnsqnqfnngcnnmwnmnbmmmslsjllbtbbpllltzzhgzztllsdllrvvhvjvbbhcbhchmchcctbcttvccgwcwpcchrcrdrdggcrrntrrfll" +
	"cffbdfflrrrgbgrbbbdqbqjbbgbgrrqwqtwqwhwghwhzwwcswsnwnqqjhjhwhfwhffdfgddgjgsjgjhgglhlwhlhssfqfhhdmdnmnppdcddfzzhm" +
	"hqqntqnnjvnjvjddcvcgcbgbbpjbjtbjjfgftgffplljfjrrhqqpddlssrvsrvrpppsllsdsqqqzzfttqsqzssjbbrnbnnrbbsrshsrrshrhwhbw" +
	"brrsrfrttfqtfqqfddvrvjrjvjsjhjsjdjqdjjlqjjjgcjcmcncfcrcwrwsrsslffzszmsszrsssrnrjjvbvpvcppptbbhhrddbcbggbqbmmsqqw" +
	"ggfpfbblmldmmpmwpwfwjfjsjnjmmpllccjzcjcwwpswshhpthhzchctcbcrrrrmvrvrdvvjmmvgmgwglghllvmllzlzzsvzzrmmhnnsjnnpvpwv" +
	"pwwmvwwdqqdffhhhmccfgfvggchcctrrmdrrhrhnhnzzgpzzgttnhthvhzzqvvvwpwqpqdppsnnrgnnhphphmhcmcrmrvvqlvqqsccqhchzhwwmvm" +
	"zmczzgsgdsggthgglrlnrlllbdllhwlwltwwcswsgssbhbsbvsbsbwbhbnncrcllttbrbppjccfpfhhgshschsccmrcmmcrrrzvvrcrggmwgwjwnw" +
	"jjbffjddjnjgngqgdgnndznndvvfqfgfvvrvqrvvpllnsszbsbdbbdzbdbzbqbzznrznzjzpptcptccvwccfscffrftrrsnsvvswvvhbhzzfbzffn" +
	"cchhcnngzzcpcmmfttsntnjjsccqbcqqmzzgppdhppdtppmffgtgvvlzlpptdtttdppqjqtqctcrrzsswwtnwtnwnqqvbbdgjhvmmzpnhfvsbddzh" +
	"gdwcnfdstvhhbzlzcfjwhlptbhmbmblprtsdmrdhbbbwpplnzgdnrzjmgzgpqbggnqvwwtntzgfwqrztqtdrsnhpfzswptggnvbszdcrmrhhtlrrf" +
	"npqrnpwrbmhlfwmdqqdbqrwbzqjbzwrgmbgrtzrhdclqfgsrtsgfwqrnnqgwsncmpgffggssrqvwjlhpsghbqdtzwmvzzvcmzsjqvprvcqwqjbcqc" +
	"qrhpwwcsrscgmfdppbgvmnrdfrppblznbstnjzwwgstjvtprjbhtpdfgrhdjnjmnlbfwggzhcngvcwvcfpcwdtdppwjrdzsnjlnrzbfqqshlnzvwsm" +
	"scgpfwjzhtwgfwgzdhbdwwzbsmfwwbmvrlrpswnjlmfbfzhwvcmgwfzssmmtjlwtrpwpwgnspbgchdncbfcpjsvtzjqtwqwjwgbhrbwvhqbcstsgsnw" +
	"sjmhrlrvzgqhqfrmnrjdrhdjwcwctpdrzctlvnfzmzwhsnfprlzgzjpqvzchlmvbhffhpfjtvsdbvbdmwgvmqpflhwwndbqthmmwshdtspsrvqdflmm" +
	"zwbqbqmpfdwjmvpbzdnqzfmhzdgldqjjvgpfcqftvjzwnzmfqdggrwlfzdhjnhmtrjbnllgqpntwmhnwtglnqdwbqdblpwnnrdwzpsqzfwqcmhqhnpsd" +
	"cwvdldphgnrtqzdbnnzdzfttldrqcztlvlrgpdqzrcthslmtqhfvbzrfgnlrprcpbsctqhspbhnjtzrzhqjzszbzdthttqmbznzssftztwlggmdqqdtf" +
	"llqjzjtvpgjfhtbwtbmtjplqnbdmsvlnqcwtdbdvfjnzgsmpnhbvvwwfbrgffjqfsccdjdwvbsdhqwfzvcpjzjbdjgrdctjplhwbdhhnbnwstvndnnwts" +
	"gbhzbvwdshvmnbwsthlrggtmddvjbfzfrnrdrqfjpslrccctzpjbwpdbhlbzfmwbltcqfngdprvfhgcszdtpnrcpdmllfnlspgrdrpwqmqbmrglvlrsmrf" +
	"qrtzzgjcvqtqzpmghjrvmdmvvqztrjzbzjwdqsmrwpqnswbzhjbzzhdvmnfdsztzdzrjssgnnfqvbtsqrrmcppjgrmnstrnrlwjvvcczqlcbmwqzdfpssfw" +
	"dfrvwtstwchdgwtrhhcmppcqlmrqlnwqccfphsdhbsbmtjvpcwwjrmrllbpnrmpbvwgwbftpdpphccwqcblcnvvbbppscmnjqgddllbnbvmmqzdffrrjtqw" +
	"llzgpqrmnlfqrptzqmdmnrfnjpvqvjbsqrhljslgqqcqqtmtbwjrpphtjgjbqpzmzrzjrfjwcdcnbsjfljclffjplnrrcfbmhphtcjrzlrvvjcznpgpnrdw" +
	"wqvgnbnzqnlcghhgwvhqbvjzfbdvhrzlqfbtlqhpltfjlfpbnjbphmmpntzqgjmwjtchwmlwvfjmfflqzpqnvvrgnbddzlfpdpdjghfgbsfddjspnfdwvqpp" +
	"ncmdgfrnvrpcrflhgjgbwdsbwblfcwbtlrrnjjdhbvmrzgsvjwgfnnhqfbvhprlmwwqgclzlbqbrdspcbhftmdscsmpwrggrmnsvjphjmzmmrlrhnmdhwjlb" +
	"mjchtvsrcplfspsssjznmzcrqnsjjtwjzvlhshbptqwwvjhjvzrhphphsbphpnzpfbwcdnqrhrrvlrwrztlpqnrcfzrncsvpzqzgslrlrwhvtgjmfncldqmvs" +
	"hlmnlpqbgvnwqfcthgrgllmqrjqmfgznspgltpptglpdcvhtzsprprbldbzhbmjsqzwvjggwhsczltcvgwqhspzpzvljwqjgrgtwswjdswlzjzslrsslvqznc" +
	"jwhbbjpbdthqpgmhfglggmlrgwdsplgscrwstntvrhjzjjlshtgmnnhvsjwfmcjbpzjcstmnpvtbgrfcfdwjljsrfhdphrdcslwhgvlnwltwchplvfzntfgcnl" +
	"svzrvnnczhhqdlwjvqprhmtjdtwmppffmszzzqtfrgnhnzqgqzhrjzgntcszstrfhhtptgvswvzvjcgcntmhzzmdgsmtgzhpfvqfnwmsjdhtfgmmbrrfsdlptc" +
	"hqqzqdqjncmtpznfssrcnmcdnthglmfzsfgltrndqsfmdftmfgchbwmzgrtjvgqtshlltthnnpqnzfrchzhdzrrnpzvfzblrmhwdwjnqdptlbvndmmlhzhvsfdl" +
	"mlhqrgqqzsdqtpczwcrwcbsftvvphfbwjrvrnrcqbbcsqgnhltwzvllljcvpwjgslbmngcdmpdvjlgcnrzwqjdgrblncpqmrgjmpqjzvdmcmwfnwqlszdgwqdfz" +
	"nhsnpsjrfwrqpqmpvhstmzgqblfmcfvwljbhdfhdmqcvrwnqcstwtzgmng"
