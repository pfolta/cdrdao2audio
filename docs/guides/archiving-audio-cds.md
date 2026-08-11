# Archiving Audio CDs

Audio CDs (**Compact Disc Digital Audio**, **CDDA** or **CD-DA**) are a standard defined in the **Red Book** of the [**Rainbow Books**](https://en.wikipedia.org/wiki/Rainbow_Books) collection of CD format specifications. This standard was formalized before the **CD-ROM** standard (**Yellow Book**). Therefore, unlike CD-ROMs, audio CDs do not contain a filesystem.

An audio CD is not made up of files such as WAV, FLAC, or MP3. Instead, the disc contains a continuous stream of digital audio sectors, along with additional metadata stored in the disc's subchannels. The track layout, indexes, gaps, and other CD-specific information are part of the physical disc structure rather than a filesystem.

Because of this, it is not possible to archive (back up) an audio CD to something like a `*.iso` file. An ISO image is a filesystem-based image designed for data discs such as CD-ROMs, DVDs, or Blu-ray media.

To correctly preserve an audio CD, the archive must capture the raw CD sectors (ideally including subchannel data) and the disc layout information.

## Using `cdrdao`

[`cdrdao`](https://cdrdao.sourceforge.net/) (**CD-R Disk-At-Once**) is a tool designed specifically for creating and writing optical disc images using the **disk-at-once (DAO)** recording method.

Unlike tools that operate on files and filesystems, `cdrdao` works at the disc layout level. It can read a CD into a pair of files:

- A **TOC file** describing the structure of the disc:
  - Track layout
  - Track types
  - Start positions
  - Gaps
  - Index markers
  - Session information
- A **BIN file** containing the raw sector data extracted from the disc

Together, these files represent a complete logical archive of the disc structure and can later be used to recreate the disc.

### Creating an archive

To create an archive including subchannel data, use:

```bash
cdrdao read-cd --read-raw --read-subchan rw_raw --datafile CD.bin CD.toc
```

> [!NOTE]
> This guide assumes the optical drive is configured as the default device. Use `--device` if multiple drives are present.

This creates a `CD.toc` file and a corresponding `CD.bin` binary data file containing the extracted disc sectors. Together, these files describe the original disc layout and contents. Both are needed to interpret the archive and to recreate the disc.

> **Command Explanation**
>
> - `cdrdao read-cd` reads a disc and creates a TOC file describing its layout and a BIN file containing its sector data.
> - `--read-raw` reads the disc's raw sectors instead of extracting only the audio data. This preserves the original sector layout required to recreate the CD accurately.
> - `--read-subchan rw_raw` reads the CD subchannel data, including information such as CD-Text and index markers where present. Using `rw_raw` preserves the complete subchannel data rather than only extracting selected information.
> - `--datafile CD.bin` specifies the filename of the BIN file to be created.
> - `CD.toc` specifies the filename of the TOC file to be created.

### Verifying an archive

After creating an archive, it is good practice to verify that the generated TOC is readable and complete:

```bash
cdrdao show-toc CD.toc
```

This displays the disc structure recorded in the TOC file.

### Recreating the original disc

A `cdrdao` archive can be written back to a blank CD-R using the `cdrdao write` command:

```bash
cdrdao write --speed 40 CD.toc
```

> **Command Explanation**
>
> - `cdrdao write` writes a disc image described by a TOC file.
> - `--speed 40` selects the writing speed. Choose a speed supported by your drive and blank CD-R. A slower speed may improve compatibility with older drives and media.
> - `CD.toc` specifies the TOC file created during the read operation.

The resulting disc should contain the same CD-DA structure as the original disc and should be playable in standard audio CD players.

## Enhanced CDs

Enhanced Music CDs (**E-CD**, also called **CD-Extra**, **CD-Plus** or **CD+**) are an extension defined in the [**Blue Book**](https://en.wikipedia.org/wiki/Blue_Book_(CD_standard)). They are multisession discs and contain both audio tracks and data tracks. The first session contains CD-DA audio tracks and can therefore be played back with regular CD players. The second session contains CD-ROM data containing files such as music videos or software and can be read by a computer with a CD-ROM drive.

Because an Enhanced CD contains multiple sessions, archiving only the audio CD portion of the disc is not sufficient. The audio session and the data session must both be preserved.

`cdrdao` operates on one session at a time. Therefore, each session has to be archived separately.

To confirm the number of sessions on a disc, use:

```bash
cdrdao disk-info
```

For an Enhanced CD, this will show something like:

```
...
Sessions             : 2
...
```

Then read each session separately:

```bash
cdrdao read-cd --session 1 --read-raw --read-subchan rw_raw --datafile session1.bin session1.toc

cdrdao read-cd --session 2 --read-raw --read-subchan rw_raw --datafile session2.bin session2.toc
```

> **Command Explanation**
>
> - `cdrdao read-cd` as above.
> - `--session #` reads session number #.

The resulting multisession archive consists of:

- `session1.toc`
- `session1.bin`
- `session2.toc`
- `session2.bin`

### Recreating the original disc

To recreate an Enhanced CD, each session must be written back in the correct order.

The CD-DA session should be written first and the disc left open (`--multi`). The data session should then be appended without `--multi` so that the disc is finalized:

```bash
cdrdao write --speed 40 --multi session1.toc

cdrdao write --speed 40 session2.toc
```

> **Command Explanation**
>
> - `cdrdao write` as above.
> - `--multi` leaves the disc open after writing so additional sessions can be added.
