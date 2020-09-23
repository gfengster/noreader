import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.PrintWriter;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.stream.Stream;

public class NOReader {

	public static void main(String[] args) throws Exception {
		String input = "/home/gfeng/Desktop/History/Original";
		String output = "/home/gfeng/Desktop/History/Readable";
		Stream<Path> stream = Files.list(Paths.get(input));

		stream.forEach(t -> {
			File in = t.toFile();
			File out = new File(output, in.getName());

			try {
				convert(in, out);
				System.out.println("Done: " + out);
			} catch (Exception e) {
				e.printStackTrace();
			}
		});
	}

	private static void convert(File in, File out) throws Exception {
		try (FileInputStream reader = new FileInputStream(in); PrintWriter writer = new PrintWriter(out)) {
			byte[] buf = new byte[1024];

			int len = 0;

			while ((len = reader.read(buf)) > -1) {
				for (int i = 0; i < len; i++) {
					byte b = buf[i] > -1 ? buf[i] : (byte) (buf[i] & 0xFF);

					if (b > 31 && b < 127 || b == 10 || b == 13) {
						writer.write(String.valueOf((char) b));
						System.out.print((char) b);
					}

				}
			}
		}
	}
}